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

	value = resolveFilters(resolveVariables(interpolateVariables(value)))

	variables[name] = value

	return name, value
}

func interpolateVariables(str string) string {
	if pos := strings.Index(str, "#{"); pos >= 0 {
		res := str[:pos]    // copy util #{
		from := str[pos+2:] // copy from #{

		if to := strings.Index(from, "}"); to > 0 {
			varName := strings.TrimSpace(from[:to])
			v := getVariable(varName)

			rest := from[to+1:]

			res += v + rest
		} else {
			// error
		}

		return interpolateVariables(res)
	}

	return str
}

func resolveVariables(str string) string {
	if pos := strings.Index(str, "$"); pos >= 0 { // got variable
		res := ""
		s := str[pos+1:] // copy from $
		res += str[:pos] // copy util $

		if n := strings.IndexAny(s, " $@#&(){}[];:,./"); n >= 0 {
			res += getVariable(s[:n])
			res += s[n:]
		} else {
			res += getVariable(s)
		}

		return resolveVariables(res)
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

func getVariableSafe(name string) string {
	value, ok := findVariable(name)

	if !ok {
		return name
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
