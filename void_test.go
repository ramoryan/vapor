package vapor

import (
	"strings"
	"testing"
)

func TestNewVoidElement(t *testing.T) {
	e := "VoidElement has been broken!"

	for _, name := range voidElements {
		v := newVoidElement(name)

		s := strings.Replace(v.render(), "\n", "", -1)

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
