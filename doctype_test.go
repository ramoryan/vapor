package vapor

import (
	"testing"
)

func TestNewDoctype(t *testing.T) {
	d, err := newDoctype("!5")
	if err != nil {
		t.Error(err)
	}

	if r, err := d.render(); r != "<!DOCTYPE html>\n" || err != nil {
		t.Error("Doctype has been broken!")
	}
}
