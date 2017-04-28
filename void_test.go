package vapor

import (
	"strings"
	"testing"
)

func TestNewVoidElement(t *testing.T) {
	e := "VoidElement has been broken!"

	for _, name := range voidElements {
		v, err := newVoidElement(name)
		if err != nil {
			t.Error(err)
		}

		r, _ := v.render()
		s := strings.Replace(r, "\n", "", -1)

		if s != "<"+name+">" {
			t.Error(e)
		}

		if !isVoidElement(name) {
			t.Error(e)
		}
	}

	if isVoidElement("div") || isVoidElement("<input>") {
		t.Error(e)
	}
}
