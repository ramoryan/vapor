package vapor

import (
	"testing"
)

func TestNewDoctype(t *testing.T) {
	d := newDoctype("!5")

	if d.render() != "<!DOCTYPE html>\n" {
		t.Error("Doctype has been broken!")
	}
}
