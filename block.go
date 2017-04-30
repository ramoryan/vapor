// block
package vapor

import (
	"reflect"
)

type block struct {
	*element
	content []string
}

func (b *forInBlock) addChild(v vaporizer) *vaporError {
	b.parent.addChild(v)
	return nil
}

func (b *block) reIndent() {
	for _, child := range b.children {
		child.setIndent(b.indent)
		child.reIndent()
	}
}

func (b *block) addContent(s string) {
	b.content = append(b.content, s)
}

func (b *block) parse() (vaporTree, *vaporError) {
	p := newParser()
	err := p.parseLines(b.content)
	if err != nil {
		return nil, err
	}

	return p.tree, nil
}

func newBlock(indent int) *block {
	c := &block{element: &element{}}
	c.indent = indent
	return c
}

func isBlockType(v vaporizer) bool {
	t := reflect.TypeOf(v).String()

	return (t == "*vapor.block")
}
