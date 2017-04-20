// variables
package vapor

import (
	"fmt"
	"strings"
)

var variables map[string]string

func isVariableInitializer(s string) bool {
	if strings.HasPrefix(s, "$") && strings.Index(s, "=") > 0 {
		return true
	}

	return false
}

func setVariable(name, value string) {
	variables[name] = value
}

func parseVariable(str string) (name, value string) {
	s := str[strings.Index(str, "$")+1:]
	index := strings.Index(s, "=")
	name = strings.TrimSpace(s[:index])
	value = strings.TrimSpace(s[index+1:])

	if variables == nil {
		variables = make(map[string]string)
	}

	value = interpolateVariables(value)

	variables[name] = value

	return name, value
}

func interpolateVariables(str string) (res string) {
	if pos := strings.Index(str, "$"); pos >= 0 { // van változó
		s := str[pos+1:] // $ jeltől lemásoljuk
		res += str[:pos] // a $ jelig másoljuk a kimenetbe

		if n := strings.IndexAny(s, " $@#&(){}[];:,./"); n >= 0 {
			res += getVariable(s[:n])
			res += s[n:]
		} else {
			res += getVariable(s)
		}

		return interpolateVariables(res)
	}

	return str
}

func getVariable(name string) string {
	value, ok := findVariable(name)

	if !ok {
		fmt.Println("variable doesn't exists:", name)
	}

	return value
}

func findVariable(name string) (string, bool) {
	if variables != nil {
		name = strings.TrimLeft(name, "$")

		if value, ok := variables[name]; ok {
			return value, true
		}
	}

	return "", false
}
