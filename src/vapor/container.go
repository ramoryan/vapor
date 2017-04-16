// container
package vapor

import (
	"reflect"
)

type container struct {
	*element
	content []string
}

func (c *container) addContent(s string) {
	c.content = append(c.content, s)
}

func (c *container) render() string {
	p := newParser()
	p.parseLines(c.content)
	return p.output
}

func newContainer(indent int) *container {
	c := &container{element: &element{}}
	c.indent = indent
	return c
}

func isContainerType(v vaporizer) bool {
	t := reflect.TypeOf(v).String()

	return (t == "*vapor.container")
}
