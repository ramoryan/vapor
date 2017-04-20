// errors
package vapor

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

type vaporError struct {
	s string
}

func (v *vaporError) Error() string {
	return v.s
}

func newVaporError(text string) *vaporError {
	return &vaporError{text}
}

func (v *vaporError) setLine(s string) {
	v.s += " " + s
}
