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
	value = unquote(value)
	a := attribute{name: name, value: interpolateVariables(value)}
	return a
}
