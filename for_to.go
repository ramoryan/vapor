// for
package vapor

import (
	"reflect"
	"strings"
)

// ---- FOR x TO Y

type forToBlock struct {
	*block
	varName string
	from    int
	to      int
}

func (f *forToBlock) parse() (vaporTree, *vaporError) {
	for i := f.from; i <= f.to; i++ {
		tree, err := f.block.parse()
		if err != nil {
			return nil, err
		}

		f.appendTree(tree) // add to self as children

		value, varErr := getVariable(f.varName)
		if varErr != nil {
			return nil, varErr
		}

		strValue := value.(string)
		intVal := strToInt(strValue, 0)
		intVal += 1

		setVariable(f.varName, intToStr(intVal, ""))
	}

	f.block.content = nil

	return nil, nil
}

func isForTo(s string) bool {
	if strings.HasPrefix(s, "for ") && strings.Index(s, " to ") > 0 {
		return true
	}

	return false
}

// TODO:
// validáció: from > to
//            to < from
func newForToBlock(s string, indent int) (*forToBlock, *vaporError) {
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

	f := &forToBlock{block: newBlock(indent)}
	f.from = strToInt(from, 0)

	// to
	toStr := strings.TrimSpace(s[toStart+2:])

	to, found := findVariable(toStr)

	if found {
		str := to.(string)

		f.to = strToInt(str, 0)
	} else {
		f.to = strToInt(toStr, 0)
	}

	f.varName = name

	return f, nil
}

func isForToBlockType(v vaporizer) bool {
	t := reflect.TypeOf(v).String()

	return (t == "*vapor.forToBlock")
}
