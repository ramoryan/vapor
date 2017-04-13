// attribute
package vapor

type attribute struct {
	name  string
	value string
}

func (a attribute) render() string {
	s := a.name

	if len(a.value) > 0 {
		s += `="` + a.value + `"`
	}

	return s
}

func newAttribute(name, value string) attribute {
	a := attribute{name: name, value: interpolateVariables(value)}
	return a
}
