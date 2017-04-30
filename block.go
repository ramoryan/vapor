// block
package vapor

import (
	"reflect"
)

type block struct {
	*element
	content []string
}

func (c *block) addContent(s string) {
	c.content = append(c.content, s)
}

func (c *block) parse() (vaporTree, *vaporError) {
	p := newParser()
	err := p.parseLines(c.content)
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
