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

func (f *forInBlock) render() (string, *vaporError) {
	res := ""

	v, err := getVariable(f.dataVarName) // findVariable(f.dataVarName)
	if err != nil {
		return "", err
	}

	// if isIterateable
	data := reflect.ValueOf(v)

	for i := 0; i < data.Len(); i++ {
		setVariable(f.iteratorVarName, intToStr(i, "")) // store the iterator actual index
		v := data.Index(i).Interface()                  // get the value

		if isStr(v) {
			val := v.(string)
			setVariable(f.valueVarName, val)
		} else if isInt(v) {
			val := v.(int)
			setVariable(f.valueVarName, intToStr(val, ""))
		}

		s, err := f.block.render()
		if err != nil {
			return "", err
		}

		res += s
	}

	return res, nil
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

func (f *forToBlock) render() (string, *vaporError) {
	res := ""

	for i := f.from; i <= f.to; i++ {
		s, err := f.block.render()
		if err != nil {
			return "", err
		}

		res += s

		value, varErr := getVariable(f.varName)
		if varErr != nil {
			return "", varErr
		}

		strValue := value.(string)
		intVal := strToInt(strValue, 0)
		intVal += 1

		setVariable(f.varName, intToStr(intVal, ""))
	}

	return res, nil
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
