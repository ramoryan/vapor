// parser
package vapor

import (
	"bufio"
	"os"
	"strings"
)

type parser struct {
	actLine    string
	actLineNum int
	parent     vaporizer
	last       vaporizer
	tree       []vaporizer
	fileName   string
	output     string
}

func (p *parser) parseFile(fileName string) *vaporError {
	p.fileName = fileName
	file, err := os.Open(fileName)

	if err != nil {
		return newVaporError(ERR_PARSER, 1, "Cannot open "+fileName)
	}

	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return p.parseLines(lines)
}

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

		// too much indent
		if p.last != nil && indent > p.last.getIndent()+8 && !isCommentType(p.last) && !isLoopBlockType(p.last) {
			return p.extendErr(newVaporError(ERR_PARSER, 3, "Too much indentation!"))
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
		} else if p.last != nil && isLoopBlockType(p.last) && indent > p.last.getIndent() {
			block := (p.last).(*loopBlock)

			if indent > p.last.getIndent() {
				// collect more rows
				block.addContent(raw)
			}

			continue
		}

		// include html or vapor file
		if isInclude(trim) {
			inc, err := include(trim)
			if err != nil {
				return p.extendErr(err)
			}

			if p.last != nil && p.last.getIndent() < indent {
				p.last.addChild(inc)
			} else {
				par := p.parent

				for par != nil && par.getIndent() >= indent {
					par = par.getParent()
				}

				par.addChild(inc)
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
		} else if isLoop(trim) { // for loop
			v, err = newLoopBlock(trim, indent) // TODO: error!
		} else if isFilter(trim) { // filter
			s, err := resolveFilters(trim)
			if err != nil {
				return p.extendErr(err)
			}

			v, err = newText(s)
		} else {
			v = resolveShortcut(trim)

			// not shortcut
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
		p.output += el.render()
	}

	return nil
}

func newParser() *parser {
	p := new(parser)
	return p
}
