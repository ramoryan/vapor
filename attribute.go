// attribute
package vapor

import (
	"strings"
)

type attribute struct {
	name  string
	value string
}

func (a *attribute) render() string {
	s := a.name

	if len(a.value) > 0 {
		s += `="` + a.value + `"`
	}

	return s
}

func parseAttribute(s string) (name, value string) {
	parts := strings.Split(s, "=")

	if len(parts) > 1 {
		return parts[0], parts[1]
	}

	return parts[0], ""
}

func isMultilineAttrCloser(s string) bool {
	if len(s) == 1 && s == ")" {
		return true
	}

	return false
}

func newAttribute(name, value string) attribute {
	if !isBetweenQuotes(value) && strings.HasPrefix(value, "$") { // it's a variable
		value = resolveVariables(value)
	} else {
		value = interpolateVariables(strings.Trim(unquote(value), "'")) // hack, ezt az unquote-nak meg kéne oldania!
	}

	name = interpolateVariables(name)

	a := attribute{name: name, value: value}
	return a
}
