package vapor

import (
	"testing"
)

func TestNewMeta(t *testing.T) {
	m := newMeta("")
	if m.render() != "<meta>\n" {
		t.Error("Meta render has been broken!")
	}
}
