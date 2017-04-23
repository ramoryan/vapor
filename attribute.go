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

func newAttribute(name, value string) (a attribute, err *vaporError) {
	if len(name) <= 0 {
		return a, newVaporError(ERR_ATTR, 1, "Name length must be greater than 0!")
	}

	if len(value) > 0 {
		if !isBetweenQuotes(value) {
			if strings.HasPrefix(value, "$") { // attr=$variable
				value, err = resolveVariables(value)
				if err != nil {
					return a, err
				}

			} else { // attr=text
				return a, newVaporError(ERR_ATTR, 2, "Quotes must be used if you want to add text!")
			}
		} else { // attr="text" or attr="#{ $text }"
			value = interpolateVariables(strings.Trim(unquote(value), "'")) // hack, ezt az unquote-nak meg kéne oldania!
		}
	}

	name = interpolateVariables(name) // #{$attr} or attr, TODO: #{   $attr }

	a = attribute{name: name, value: value}
	return a, err
}
