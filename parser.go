// parser
package vapor

import (
	"bufio"
	"os"
	"path"
	"strings"
)

type vaporTree []vaporizer

var renderCount int

type parser struct {
	actLine    string
	actLineNum int
	parent     vaporizer
	last       vaporizer
	tree       vaporTree
	fileName   string
	output     string
}

func (p *parser) parseFile(fileName string) *vaporError {
	// Check file extension, only .vapr or .html is allowed.
	ext := path.Ext(fileName)
	if ext != ".vapr" && ext != ".html" {
		return newVaporError(ERR_PARSER, 1, "File extension must be .vapr or .html! "+fileName+" given.")
	}

	// Try to open a file
	p.fileName = fileName
	file, err := os.Open(fileName)
	if err != nil {
		return newVaporError(ERR_PARSER, 2, "Cannot open "+fileName)
	}

	defer file.Close()

	// Collect the lines.
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Returns vaporError or nil.
	return p.parseLines(lines)
}

// Extends the vaporError with the most important datas:
// - actual line
// - filename
// - actual line number
func (p *parser) extendErr(err *vaporError) *vaporError {
	return err.addErrorLine(p.actLine).addFileName(p.fileName).addErrLineNum(p.actLineNum)
}

func (p *parser) parseLines(lines []string) *vaporError {
	var err *vaporError
	firstIndent := -1

	// trough the lines
	for _, raw := range lines {
		p.actLineNum++ // for error
		raw = removeComment(raw)
		trim := strings.TrimSpace(raw) // trimmelt text
		p.actLine = trim               // for error
		indent := calcIndent(raw)

		// empty line
		if len(trim) <= 0 {
			continue
		}

		// too much indent compared the last
		if p.last != nil && indent > p.last.getIndent()+8 &&
			!isCommentType(p.last) && !isForToBlockType(p.last) && !isForInBlockType(p.last) {
			return p.extendErr(newVaporError(ERR_PARSER, 3, "Too much indentation!"))
		} else if firstIndent == -1 && indent > 0 {
			// return p.extendErr(newVaporError(ERR_PARSER, 3, "Too much indentation!"))
		}

		// new variable
		if isVariableInitializer(trim) {
			_, _, err := parseVariable(trim)
			if err != nil {
				return p.extendErr(err)
			}

			continue
		}

		// is collecting multiline attributes?
		if p.last != nil && p.last.needMultilineAttrs() {
			parIndent := p.last.getIndent()

			// we want to close it
			if isMultilineAttrCloser(trim) && indent == parIndent {
				p.last.closeMultilineAttrs()
			} else { // adds a new attr
				if indent == parIndent+8 {
					name, value, err := parseAttribute(trim)
					if err != nil {
						return p.extendErr(err)
					}

					p.last.addAttr(name, value)
				} else {
					return newVaporError(ERR_PARSER, 2, "Syntax error! You have to close the multiline attributes before: "+trim).
						addFileName(p.fileName).
						addErrLineNum(p.actLineNum)
				}
			}

			continue
		}

		// comment
		if p.last != nil && isCommentType(p.last) && indent >= p.last.getIndent()+8 {
			(p.last).(*comment).addContent(raw) // cast last one to comment and add this line as content
			continue
			// loop
			// --- for to
		} else if p.last != nil && isForToBlockType(p.last) && indent > p.last.getIndent() {
			block := (p.last).(*forToBlock)

			// collect more rows
			block.addContent(raw)

			continue
			// --- for in
		} else if p.last != nil && isForInBlockType(p.last) && indent > p.last.getIndent() {
			block := (p.last).(*forInBlock)

			// collect more rows
			block.addContent(raw)

			continue
		}

		// include html or vapor file
		if isInclude(trim) {
			tree, err := include(trim)
			if err != nil {
				return p.extendErr(err)
			}

			if p.last != nil && p.last.getIndent() < indent {
				p.last.appendTree(tree)
			} else {
				par := p.parent

				for par != nil && par.getIndent() >= indent {
					par = par.getParent()
				}

				par.appendTree(tree)
			}

			continue
		}

		// html element or shortcut
		var v vaporizer

		if strings.HasPrefix(trim, "!5") {
			v, err = newDoctype(raw)
		} else if strings.HasPrefix(trim, "html") {
			v, err = newHtml(raw)
		} else if strings.HasPrefix(trim, "head") && !strings.HasPrefix(trim, "header") {
			v, err = newHead(raw)
		} else if strings.HasPrefix(trim, "meta") {
			v, err = newMeta(raw)
		} else if isText(trim) { // text
			v, err = newText(raw)
		} else if isVoidElement(trim) { // void / selfclosed element?
			v, err = newVoidElement(raw)
		} else if isComment(trim) { // comment
			v = newComment(raw)
		} else if isForTo(trim) { // for x to y
			v, err = newForToBlock(trim, indent) // TODO: error!
		} else if isForIn(trim) { // for i, v in data
			v, err = newForInBlock(trim, indent)
		} else if isFilter(trim) { // filter
			s, err := resolveFilters(trim)
			if err != nil {
				return p.extendErr(err)
			}

			v, err = newText(s)
		} else {
			v = resolveShortcut(trim)
			if v == nil {
				v, err = newElement(raw)
			}
		}

		if err != nil {
			return p.extendErr(err)
		}

		if firstIndent == -1 { // first parsable line
			firstIndent = indent
		}

		if indent == firstIndent { // new branch of the tree
			p.tree = append(p.tree, v)
		} else if indent > p.last.getIndent() { // new parent
			v.setParent(p.last)
			p.parent = v
			p.last.addChild(v)
		} else if indent == p.last.getIndent() { // the last parent needed
			v.setParent(p.last.getParent())
			p.last.getParent().addChild(v)
		} else if indent < p.last.getIndent() { // searching for last parent by indent
			for p.parent != nil && p.parent.getIndent() >= indent {
				p.parent = p.parent.getParent()
			}

			v.setParent(p.parent)
			p.parent.addChild(v)
		}

		p.last = v
	}

	for _, el := range p.tree {
		_, err := el.parse()
		if err != nil {
			return p.extendErr(err)
		}
	}

	for _, el := range p.tree {
		el.reIndent()
		out, err := el.render()
		if err != nil {
			return p.extendErr(err)
		}

		p.output += out
	}

	return nil
}

func newParser() *parser {
	p := new(parser)
	return p
}
