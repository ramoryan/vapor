// errors
package vapor

import (
	"fmt"
)

const (
	ERR_PARSER   = 1
	ERR_ELEMENT  = 2
	ERR_ATTR     = 3
	ERR_INCLUDE  = 4
	ERR_FILTER   = 5
	ERR_VARIABLE = 6
	ERR_HEAD     = 7
	ERR_LOOP     = 8
)

type vaporError struct {
	fileName   string
	errLineNum int
	errType    int
	errNum     int
	errLine    string
	errMsg     string
}

func (v *vaporError) String() string {
	// return intToStr(v.errType, "") + intToStr(v.errNum, "") + " " + v.errMsg + "\n" + v.errMsg
	zero := ""

	if v.errNum < 10 {
		zero = "0"
	}

	s := fmt.Sprintf("Error: %v%v%v - %v\n", v.errType, zero, v.errNum, getErrorType(v.errType))
	s += v.errMsg + "\n"

	if len(v.fileName) > 0 {
		s += v.fileName + ":" + intToStr(v.errLineNum, "") + "\n" + v.errLine
	}

	return s
}

func (v *vaporError) addErrLineNum(n int) *vaporError {
	if v.errLineNum <= 0 {
		v.errLineNum = n
	}

	return v
}

func (v *vaporError) addFileName(s string) *vaporError {
	if len(v.fileName) == 0 {
		v.fileName = s
	}

	return v
}

func (v *vaporError) addErrorLine(s string) *vaporError {
	if len(v.errLine) == 0 {
		v.errLine = s
	}

	return v
}

func getErrorType(n int) string {
	switch n {
	case ERR_PARSER:
		return "Parser"
	case ERR_ELEMENT:
		return "Element"
	case ERR_ATTR:
		return "Attribute"
	case ERR_INCLUDE:
		return "Include"
	case ERR_FILTER:
		return "Filter"
	case ERR_VARIABLE:
		return "Variable"
	case ERR_LOOP:
		return "Loop"
	}

	return "Unknown"
}

func newVaporError(errType, errNum int, errMsg string) *vaporError {
	e := &vaporError{errType: errType, errNum: errNum, errMsg: errMsg}
	return e
}
