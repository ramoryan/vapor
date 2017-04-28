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
	txt, err := newText("| Text")
	if err != nil {
		t.Error(err)
	}

	if r, err := txt.render(); r != "Text\n" || err != nil {
		t.Error("Text rendering is broken!")
	}

	txt, err = newText("| text and $21 in cash")
	if err != nil {
		t.Error(err)
	}

	if r, err := txt.render(); r != "text and $21 in cash\n" || err != nil {
		t.Error("Text rendering with $ sign is broken!")
	}

	clearVariables()
	txt, err = newText("|     text and #{ $undefined }")
	if err == nil {
		t.Error("newText must returns undefined variable error!")
	}

	setVariable("interpolateMe", "Say good bye!")
	txt, err = newText("| text text #{   $interpolateMe    } text")
	if err != nil {
		t.Error(err)
	}

	if r, err := txt.render(); r != "text text Say good bye! text\n" || err != nil {
		t.Error("Variable interpolation in text has been broken!")
	}
}
