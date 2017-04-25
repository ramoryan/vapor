package vapor

import (
	"strings"
	"unicode"
)

type attrMap map[string]attribute

type element struct {
	attributes    attrMap
	children      []vaporizer
	parent        vaporizer
	name          string
	indent        int
	raw           string
	isVoid        bool
	attrFields    []string
	inlineText    string
	needMoreAttrs bool
}

func (e *element) render() string {
	spc := renderIndent(e.indent)

	s := spc + "<" + e.name

	for _, attr := range e.attributes {
		s += " " + attr.render()
	}

	s += ">"

	hasText := len(e.inlineText) > 0
	hasChildren := len(e.children) > 0

	if hasText || hasChildren || e.isVoid {
		s += "\n"

		if hasText {
			s += renderIndent(e.indent+8) + e.inlineText + "\n"
		}

		for _, child := range e.children {
			s += child.render()
		}
	}

	if !e.isVoid {
		if hasText || hasChildren {
			s += spc + "</" + e.name + ">\n"
		} else {
			s += "</" + e.name + ">\n"
		}
	}

	return s
}

func (e *element) getIndent() int {
	return e.indent
}

func (e *element) setParent(v vaporizer) {
	e.parent = v
}

func (e *element) addChild(v vaporizer) *vaporError {
	e.children = append(e.children, v)
	v.setIndent(e.indent + 8)
	return nil
}

func (e *element) getAttributes() attrMap {
	return e.attributes
}

func (e *element) getChildren() []vaporizer {
	return e.children
}

func (e *element) addAttr(name, value string) *vaporError {
	a, err := newAttribute(name, value)
	if err != nil {
		return err
	}

	if e.attributes == nil {
		e.attributes = make(attrMap)
	} else if _, exists := e.attributes["foo"]; exists {
		return newVaporError(ERR_ELEMENT, 5, "Attribute with this name has been already initialized: "+name)
	}

	e.attributes[name] = a

	return nil
}

func (e *element) addAttrQ(name, value string) *vaporError {
	return e.addAttr(name, `"`+value+`"`)
}

func (e *element) setIndent(indent int) {
	e.indent = indent
}

func (e *element) getParent() vaporizer {
	return e.parent
}

func (e *element) getName() string {
	return e.name
}

func (e *element) splitToFields() {
	s := strings.TrimSpace(e.raw)
	attrStart := strings.Index(s, "(")
	attrEnd := strings.LastIndex(s, ")")

	if attrStart > 0 && attrEnd > 0 {
		attrs := s[attrStart+1 : attrEnd]
		e.name = s[:attrStart]

		lastQuote := rune(0)
		f := func(c rune) bool {
			switch {
			case c == lastQuote:
				lastQuote = rune(0)
				return false
			case lastQuote != rune(0):
				return false
			case unicode.In(c, unicode.Quotation_Mark):
				lastQuote = c
				return false
			default:
				return unicode.IsSpace(c)
			}
		}

		e.attrFields = strings.FieldsFunc(attrs, f)

		if len(e.attrFields) <= 0 {
			e.attrFields = append(e.attrFields, attrs)
		}

		if len(s) > attrEnd {
			e.inlineText = s[attrEnd+1:]
		}

	} else {
		spaceIndex := strings.Index(s, " ")

		if spaceIndex > 0 {
			e.name = s[:spaceIndex]
			e.inlineText = s[spaceIndex:]
		} else {
			e.name = s
		}
	}

	// comment lehetőség a taggel egy sorban
	if len(e.inlineText) > 0 {
		if pos := strings.Index(e.inlineText, "//"); pos >= 0 {
			e.inlineText = strings.TrimSpace(e.inlineText[:pos])
		}
	}
}

func (e *element) hasAttr(attrName, attrValue string) bool {
	for _, a := range e.getAttributes() {
		if a.name == attrName && a.value == attrValue {
			return true
		}
	}

	return false
}

func (e *element) setAttributes() *vaporError {
	attrs := e.attrFields

	if len(attrs) > 0 {
		for i := 0; i < len(attrs); i++ {
			field := attrs[i]

			name, value, err := parseAttribute(field)

			if err != nil {
				return err
			}

			err = e.addAttr(name, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// TODO: fkin great refact needed!
func (e *element) resolveShortCuts(name string) (err *vaporError) {
	if strings.HasPrefix(name, "#") {
		id := name[1:] // remove #

		if pos := strings.Index(id, "."); pos > 0 {
			class := id[pos+1:]
			id = id[:pos]

			if len(class) <= 0 {
				return newVaporError(ERR_ELEMENT, 7, "Class shortcut has no value!")
			}

			err = e.addAttrQ("class", class)
			if err != nil {
				return err
			}
		}

		if len(id) <= 0 {
			return newVaporError(ERR_ELEMENT, 8, "Id shortcut has no value!")
		}

		err = e.addAttrQ("id", id)
		if err != nil {
			return err
		}

		e.name = "div"
	} else if strings.HasPrefix(name, ".") {
		class := name[1:] // remove .

		if pos := strings.Index(class, "#"); pos > 0 {
			id := class[pos+1:]
			class = class[:pos]

			if len(id) <= 0 {
				return newVaporError(ERR_ELEMENT, 8, "Id shortcut has no value!")
			}

			err = e.addAttrQ("id", id)
			if err != nil {
				return err
			}
		}

		if len(class) <= 0 {
			return newVaporError(ERR_ELEMENT, 7, "Class shortcut has no value!")
		}

		err = e.addAttrQ("class", class)
		if err != nil {
			return err
		}

		e.name = "div"
	} else if pos := strings.Index(name, "#"); pos > 0 {
		id := name[pos+1:]
		name = name[:pos]

		if classPos := strings.Index(id, "."); classPos > 0 {
			class := id[classPos+1:]
			id = id[:classPos]

			if len(class) <= 0 {
				return newVaporError(ERR_ELEMENT, 7, "Class shortcut has no value!")
			}

			err = e.addAttrQ("class", class)
			if err != nil {
				return err
			}
		}

		if len(id) <= 0 {
			return newVaporError(ERR_ELEMENT, 8, "Id shortcut has no value!")
		}

		err = e.addAttrQ("id", id)
		if err != nil {
			return err
		}

		e.name = name

		err := e.resolveShortCuts(name)
		if err != nil {
			return err
		}
	} else if pos := strings.Index(name, "."); pos > 0 {
		class := name[pos+1:]
		name = name[:pos]

		if idPos := strings.Index(class, "#"); idPos > 0 {
			id := class[idPos+1:]
			class = class[:idPos]

			if len(id) <= 0 {
				return newVaporError(ERR_ELEMENT, 8, "Id shortcut has no value!")
			}

			err = e.addAttrQ("id", id)
			if err != nil {
				return err
			}
		}

		if len(class) <= 0 {
			return newVaporError(ERR_ELEMENT, 7, "Class shortcut has no value!")
		}

		err = e.addAttrQ("class", class)
		if err != nil {
			return err
		}

		e.name = name

		err = e.resolveShortCuts(name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *element) resolveMultilineAttrOpener() {
	s := e.name

	if strings.Contains(s, "(") && !strings.Contains(s, ")") { // nyitni akar multilinet?
		if strings.Contains(s, " ") || !strings.HasSuffix(s, "(") {
			// parse error
		} else {
			e.needMoreAttrs = true
			s = strings.Replace(s, "(", "", -1)
		}
	}

	e.name = s
}

func (e *element) needMultilineAttrs() bool {
	return e.needMoreAttrs
}

func (e *element) closeMultilineAttrs() {
	e.needMoreAttrs = false
}

func newElement(raw string) (e *element, err *vaporError) {
	e = &element{raw: raw}
	e.indent = calcIndent(raw)
	e.splitToFields()

	err = e.setAttributes()
	if err != nil {
		return e, err
	}

	e.resolveMultilineAttrOpener()

	err = e.resolveShortCuts(e.name)
	if err != nil {
		return e, err
	}

	return e, nil
}
