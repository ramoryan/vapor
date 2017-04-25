package vapor

import (
	"testing"
)

func TestNewDoctype(t *testing.T) {
	d, err := newDoctype("!5")
	if err != nil {
		t.Error(err)
	}

	if d.render() != "<!DOCTYPE html>\n" {
		t.Error("Doctype has been broken!")
	}
}
