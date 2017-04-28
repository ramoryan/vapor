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

func (c *block) render() (string, *vaporError) {
	p := newParser()
	err := p.parseLines(c.content)
	if err != nil {
		return "", err
	}

	return p.output, nil
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
