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

	// get the data by varName
	v, err := getVariable(f.dataVarName)
	if err != nil {
		return "", err
	}

	// is Map, Slice or String?
	if !isIterateable(v) {
		return "", newVaporError(ERR_LOOP, 5, "Data must be Map, Slice or String!")
	}

	// string
	if isStr(v) {
		for index, value := range v.(string) {
			setVariable(f.iteratorVarName, intToStr(index, "")) // store the iterator actual index
			setVariable(f.valueVarName, string(value))

			s, err := f.block.render()
			if err != nil {
				return "", err
			}

			res += s
		}
	} else if isMap(v) {
		data := reflect.ValueOf(v)

		for _, key := range data.MapKeys() {
			setVariable(f.iteratorVarName, key.String())

			value := data.MapIndex(key).Interface() // get the value

			if isStr(value) {
				v := value.(string)
				setVariable(f.valueVarName, v)
			} else if isInt(value) {
				v := value.(int)
				setVariable(f.valueVarName, intToStr(v, ""))
			}

			s, err := f.block.render()
			if err != nil {
				return "", err
			}

			res += s
		}
	} else { // slice
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
			} else {
				return "", newVaporError(ERR_LOOP, 6, "Not applicable type! "+typeof(v).String())
			}

			s, err := f.block.render()
			if err != nil {
				return "", err
			}

			res += s
		}
	}

	return res, nil
}

// returns it's a "for in" loop or not
func isForIn(s string) bool {
	if strings.HasPrefix(s, "for ") && strings.Index(s, " in ") > 0 {
		return true
	}

	return false
}

func newForInBlock(s string, indent int) (*forInBlock, *vaporError) {
	if !strings.HasPrefix(s, "for ") {
		return nil, newVaporError(ERR_LOOP, 1, "Loop must be start with 'for'!")
	}

	s = strings.TrimSpace(strings.TrimLeft(s, "for")) // the string without "for"

	toStart := strings.Index(s, "in")
	if toStart < 0 {
		return nil, newVaporError(ERR_LOOP, 2, "Loop must contains 'in' keyword!")
	}

	// collect the "key" and "value" initializers
	varsStr := strings.TrimSpace(s[:toStart])
	if len(varsStr) <= 0 {
		return nil, newVaporError(ERR_LOOP, 3, "Loop must contains 'key' and 'value' initializers!")
	}

	//vars := strings.Split(varsStr, ",")
	vars := splitAndTrim(varsStr, ",")
	if len(vars) != 2 {
		return nil, newVaporError(ERR_LOOP, 3, "Loop must contains 'key' and 'value' initializers!")
	}

	if !strings.HasPrefix(vars[0], "$") || !strings.HasPrefix(vars[1], "$") {
		return nil, newVaporError(ERR_LOOP, 4, "Not valid variable initializers! Use the $ sign!")
	}

	// collect the "data" variable
	dataVar := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(s[toStart:]), "in"))

	_, err := getVariable(dataVar)
	if err != nil {
		return nil, err
	}

	f := &forInBlock{block: newBlock(indent)}
	f.iteratorVarName = strings.TrimLeft(vars[0], "$")
	f.valueVarName = strings.TrimLeft(vars[1], "$")
	f.dataVarName = dataVar

	return f, nil
}

func isForInBlockType(v vaporizer) bool {
	t := reflect.TypeOf(v).String()

	return (t == "*vapor.forInBlock")
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
