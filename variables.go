// variables
package vapor

import (
	"strings"
)

var variables map[string]interface{}

func isVariableInitializer(s string) bool {
	if strings.HasPrefix(s, "$") && strings.Index(s, "=") > 0 {
		return true
	}

	return false
}

func setVariable(name string, value interface{}) {
	if variables == nil {
		variables = make(map[string]interface{})
	}

	variables[name] = value
}

func parseVariable(str string) (name, value string, e *vaporError) {
	if !strings.HasPrefix(str, "$") {
		return "", "", newVaporError(ERR_VARIABLE, 1, "$ sign must be the first character!")
	}

	if strings.Index(str, "=") < 0 {
		return "", "", newVaporError(ERR_VARIABLE, 2, "You have to use equation (=) sign!")
	}

	s := str[strings.Index(str, "$")+1:]
	index := strings.Index(s, "=")
	name = strings.TrimSpace(s[:index])
	value = strings.TrimSpace(s[index+1:])

	if len(value) == 0 {
		return "", "", newVaporError(ERR_VARIABLE, 5, "You have to define variable value!")
	}

	var err *vaporError

	value, err = interpolateVariables(value)
	if err != nil {
		return "", "", err
	}

	value, err = resolveVariables(value)
	if err != nil {
		return "", "", err
	}

	value, err = resolveFilters(value)
	if err != nil {
		return "", "", err
	}

	setVariable(name, value)

	return name, value, nil
}

func interpolateVariables(str string) (string, *vaporError) {
	if pos := strings.Index(str, "#{"); pos >= 0 {
		res := str[:pos]    // copy util #{
		from := str[pos+2:] // copy from #{

		if to := strings.Index(from, "}"); to > 0 {
			varName := strings.TrimSpace(from[:to])
			v, err := getVariable(varName)
			if err != nil {
				return "", err
			}

			rest := from[to+1:]
			res += v.(string) + rest
		} else {
			return "", newVaporError(ERR_VARIABLE, 4, "You have to close variable interpoltion!")
		}

		return interpolateVariables(res)
	}

	return str, nil
}

func resolveVariables(str string) (string, *vaporError) {
	if pos := strings.Index(str, "$"); pos >= 0 { // got variable
		res := ""
		s := str[pos+1:] // copy from $
		res += str[:pos] // copy util $

		if n := strings.IndexAny(s, " $@#&(){}[];:,./"); n >= 0 {
			val, err := getVariable(s[:n])
			if err != nil {
				return "", err
			}

			res += val.(string)
			res += s[n:]
		} else {
			val, err := getVariable(s)
			if err != nil {
				return "", err
			}

			res += val.(string)
		}

		r, err := resolveVariables(res)
		if err != nil {
			return "", err
		}

		return r, nil
	}

	return str, nil
}

func getVariable(name string) (interface{}, *vaporError) {
	value, ok := findVariable(name)

	if !ok {
		return "", newVaporError(ERR_VARIABLE, 3, "Variable doesn't exists: "+name)
	}

	return value, nil
}

func getVariableSafe(name string) interface{} {
	value, ok := findVariable(name)

	if !ok {
		return name
	}

	return value
}

func findVariable(name string) (interface{}, bool) {
	if variables != nil {
		name = strings.TrimLeft(name, "$")

		if value, ok := variables[name]; ok {
			return value, true
		}
	}

	return "", false
}

func clearVariables() {
	for k := range variables {
		delete(variables, k)
	}
}
