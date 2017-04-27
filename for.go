// for
package vapor

import (
	"reflect"
	"strings"
)

// ---- FOR $i, $v IN $slice|$map
type forInBlock struct {
	*block
	iteratorVarName string
	valueVarName    string
	dataVarName     string
	iterator        interface{} // string or int
	value           interface{} // anything
}

func (f *forInBlock) render() string {
	res := ""

	data, _ := findVariable(f.dataVarName)
	slice := reflect.ValueOf(data)

	for i := 0; i < slice.Len(); i++ {
		setVariable(f.iteratorVarName, intToStr(i, ""))

		v := slice.Index(i)
		val := v.Interface().(string)
		// setVariable(f.valueVarName, intToStr(slice.Index(i), ""))
		setVariable(f.valueVarName, val)

		s := f.block.render()
		res += s
	}

	return res
}

func isForIn(s string) bool {
	if strings.HasPrefix(s, "for ") && strings.Index(s, " in ") > 0 {
		return true
	}

	return false
}

func newForInBlock(s string, indent int) (*forInBlock, *vaporError) {
	if strings.Index(s, "for ") < 0 {
		return nil, newVaporError(ERR_LOOP, 1, "Loop must be start with 'for'!")
	}

	s = strings.TrimSpace(strings.TrimLeft(s, "for"))

	toStart := strings.Index(s, "in")
	if toStart < 0 {
		return nil, newVaporError(ERR_LOOP, 2, "Loop must contains 'in' keyword!")
	}

	f := &forInBlock{block: newBlock(indent)}
	f.iteratorVarName = "i"
	f.valueVarName = "v"
	f.dataVarName = "vaporSlice"

	return f, nil
}

// ---- FOR x TO Y

type forToBlock struct {
	*block
	varName string
	from    int
	to      int
}

func (f *forToBlock) render() string {
	res := ""

	for i := f.from; i <= f.to; i++ {
		s := f.block.render()
		res += s

		value, _ := getVariable(f.varName)
		strValue := value.(string)
		intVal := strToInt(strValue, 0)
		intVal += 1

		setVariable(f.varName, intToStr(intVal, ""))
	}

	return res
}

func isForTo(s string) bool {
	if strings.HasPrefix(s, "for ") && strings.Index(s, " to ") > 0 {
		return true
	}

	return false
}

func isForInBlockType(v vaporizer) bool {
	t := reflect.TypeOf(v).String()

	return (t == "*vapor.forInBlock")
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
