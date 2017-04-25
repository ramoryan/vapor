package vapor

import (
	"testing"
)

func TestNewMeta(t *testing.T) {
	m, err := newMeta("")
	if err != nil {
		t.Error(err)
	}

	if m.render() != "<meta>\n" {
		t.Error("Meta render has been broken!")
	}
}
