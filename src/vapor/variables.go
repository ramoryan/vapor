// variables
package vapor

import (
	"fmt"
	"strings"
)

var variables map[string]string

func isVariableInitializer(s string) bool {
	if strings.HasPrefix(s, "$") && strings.Index(s, ":") > 0 {
		return true
	}

	return false
}

func parseVariable(str string) {
	s := str[strings.Index(str, "$")+1:]
	name := strings.TrimSpace(s[:strings.Index(s, ":")])
	value := strings.TrimSpace(s[strings.Index(s, ":")+1:])

	if variables == nil {
		variables = make(map[string]string)
	}

	variables[name] = interpolateVariables(value)
}

func interpolateVariables(str string) (res string) {
	if strings.Contains(str, "$") {
		vars := strings.Split(str, "$")
		for i, v := range vars {
			if i > 0 {
				n := strings.IndexAny(v, " @#&(){}[];:,./")
				if n >= 0 {
					res += getVariable(v[:n])
					res += v[n:]
				} else {
					res += getVariable(v)
				}
			} else {
				res += v
			}
		}
		return
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
		if value, ok := variables[name]; ok {
			return value, true
		}
	}

	return "", false
}
