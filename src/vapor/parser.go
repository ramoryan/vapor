// parser
package vapor

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type parser struct {
	parent vaporizer
	last   vaporizer
	tree   []vaporizer
	output string
}

func (p *parser) parseFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	p.parseLines(lines)
}

func (p *parser) parseLines(lines []string) {
	// végig a sorokon
	for _, raw := range lines {
		trim := strings.TrimSpace(raw) // trimmelt text
		indent := calcIndent(raw)

		// üres sor
		if len(trim) <= 0 {
			continue
		}

		// új változó
		if isVariableInitializer(trim) {
			parseVariable(trim)
			continue
		}

		// multiline attribútumokat vár
		if p.last != nil && p.last.needMultilineAttrs() {
			parIndent := p.last.getIndent()

			// le akarjuk zárni
			if isMultilineAttrCloser(trim) && indent == parIndent {
				p.last.closeMultilineAttrs()
			} else { // újat akarunk hozzáadni
				if indent == parIndent+8 {
					p.last.addAttr(parseAttribute(trim))
				} else {
					// indent error
				}
			}

			continue
		}

		// valamilyen elem vagy comment
		var v vaporizer

		if strings.HasPrefix(trim, "!5") {
			v = newDoctype(raw)
		} else if strings.HasPrefix(trim, "html") {
			v = newHtml(raw)
		} else if strings.HasPrefix(trim, "head") && !strings.HasPrefix(trim, "header") {
			v = newHead(raw)
		} else if strings.HasPrefix(trim, "meta") {
			v = newMeta(raw)
		} else if isText(trim) {
			v = newText(raw)
		} else if isVoidElement(trim) { // void / selfclosed element?
			v = newVoidElement(raw)
		} else if isComment(trim) {
			v = newComment(raw)
		} else if isLoop(trim) {
			v = handleLoop(trim, indent)
		} else if p.last != nil && isCommentType(p.last) && indent >= p.last.getIndent()+8 {
			(p.last).(*comment).addContent(raw) // az előzőt castoljuk commentté és hozzáadjuk a szöveget, mint tartalom
			continue
		} else if p.last != nil && isContainerType(p.last) {
			container := (p.last).(*container)

			if indent > p.last.getIndent() {
				// gyűjtjük a sorokat
				container.addContent(raw)
			} else {
				// végrehajtjuk
			}

			continue
		} else {
			v = newElement(raw)
		}

		if indent <= 0 { // megy a fa következő ágának
			p.parent = nil
			p.tree = append(p.tree, v)
		} else if indent > p.last.getIndent() { // ez az új parent
			v.setParent(p.last)
			p.parent = v
			p.last.addChild(v)
		} else if indent == p.last.getIndent() { // a legútobbi parent kell
			v.setParent(p.last.getParent())
			p.last.getParent().addChild(v)
		} else if indent < p.last.getIndent() { // megkeressük az indent alapján ezt megelőző parentet
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
}

func newParser() *parser {
	p := new(parser)
	return p
}
