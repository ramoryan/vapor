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

		strValue, _ := getVariable(l.varName)
		intVal := strToInt(strValue, 0)
		intVal += 1

		setVariable(l.varName, intToStr(intVal, ""))
	}

	return res
}

func isLoop(s string) bool {
	if strings.HasPrefix(s, "for ") && strings.Index(s, " to ") > 0 {
		return true
	}

	return false
}

// TODO:
// validáció: from > to
//            to < from
func newLoopBlock(s string, indent int) (*loopBlock, *vaporError) {
	if strings.Index(s, "for ") < 0 {
		return nil, newVaporError(ERR_LOOP, 1, "Loop must be start with 'for'!")
	}

	s = strings.TrimSpace(strings.TrimLeft(s, "for"))

	toStart := strings.Index(s, "to")

	if toStart < 0 {
		return nil, newVaporError(ERR_LOOP, 2, "Loop must contains 'to' keyword!")
	}

	// from
	fromStr := strings.TrimSpace(s[:toStart])
	name, from, err := parseVariable(fromStr)
	if err != nil {
		return nil, err
	}

	l := &loopBlock{block: newBlock(indent)}
	l.from = strToInt(from, 0)

	// to
	toStr := strings.TrimSpace(s[toStart+2:])

	to, found := findVariable(toStr)

	if found {
		l.to = strToInt(to, 0)
	} else {
		l.to = strToInt(toStr, 0)
	}

	l.varName = name

	return l, nil
}

func isLoopBlockType(v vaporizer) bool {
	t := reflect.TypeOf(v).String()

	return (t == "*vapor.loopBlock")
}
