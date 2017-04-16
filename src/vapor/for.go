// for
package vapor

import (
	"reflect"
	"strings"
)

// for 1 to 4

type loopContainer struct {
	*container
	from int
	to   int
}

func (l *loopContainer) render() string {
	s := l.container.render()
	res := ""

	for i := l.from; i <= l.to; i++ {
		res += s
	}

	return res
}

func isLoop(s string) bool {
	if strings.HasPrefix(s, "for") && strings.Index(s, " to ") > 0 {
		return true
	}

	return false
}

func newLoopContainer(s string, indent int) *loopContainer {
	f := strings.Fields(s)

	c := &loopContainer{container: newContainer(indent)}
	c.from = strToInt(f[1], 0)
	c.to = strToInt(f[3], 0)

	return c
}

func isLoopContainerType(v vaporizer) bool {
	t := reflect.TypeOf(v).String()

	return (t == "*vapor.loopContainer")
}
