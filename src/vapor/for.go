// for
package vapor

import (
	"reflect"
	"strings"
)

// for 1 to 4

type loopBlock struct {
	*block
	varName string
	from    int
	to      int
}

func (l *loopBlock) render() string {
	res := ""

	for i := l.from; i <= l.to; i++ {
		s := l.block.render()
		res += s

		val := strToInt(getVariable(l.varName), 0)
		val += 1

		setVariable(l.varName, intToStr(val, ""))
	}

	return res
}

func isLoop(s string) bool {
	if strings.HasPrefix(s, "for") && strings.Index(s, " to ") > 0 {
		return true
	}

	return false
}

// TODO:
// validáció: from > to
//            to < from
func newLoopBlock(s string, indent int) *loopBlock {
	f := strings.Fields(s)

	name, from := parseVariable(f[1])

	l := &loopBlock{block: newBlock(indent)}
	l.from = strToInt(from, 0)

	to, found := findVariable(f[3])

	if found {
		l.to = strToInt(to, 0)
	} else {
		l.to = strToInt(f[3], 0)
	}

	l.varName = name

	return l
}

func isLoopBlockType(v vaporizer) bool {
	t := reflect.TypeOf(v).String()

	return (t == "*vapor.loopBlock")
}
