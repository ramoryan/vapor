package vapor

import (
	"testing"
)

func TestNewMeta(t *testing.T) {
	m, err := newMeta("")
	if err != nil {
		t.Error(err)
	}

	if r, err := m.render(); r != "<meta>\n" || err != nil {
		t.Error("Meta render has been broken!")
	}
}
