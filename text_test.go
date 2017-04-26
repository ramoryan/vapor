package vapor

import (
	"testing"
)

func TestIsText(t *testing.T) {
	if !isText("| Rock and roll shoes") {
		t.Error("String starts with pipe (|) must be a vapor text!")
	}
}

func TestNewText(t *testing.T) {
	txt, _ := newText("| Text")
	if txt.render() != "Text\n" {
		t.Error("Text rendering is broken!")
	}

	txt, _ = newText("| text and $21 in cash")
	if txt.render() != "text and $21 in cash\n" {
		t.Error("Text rendering with $ sign is broken!")
	}

	clearVariables()
	var err *vaporError
	txt, err = newText("|     text and #{ $undefined }")
	if err == nil {
		t.Error("newText must returns undefined variable error!")
	}

	setVariable("interpolateMe", "Say good bye!")
	txt, err = newText("| text text #{   $interpolateMe    } text")
	if err != nil || txt.render() != "text text Say good bye! text\n" {
		t.Error("Variable interpolation in text has been broken!")
	}
}
