package vapor

import (
	"strings"
	"testing"
)

func TestNewComment(t *testing.T) {
	clearVariables()

	// vapor comment
	c := newComment("// vapor comment")
	if c.commentType != C_VAPOR {
		t.Error("It must be a vapor comment!")
	}

	if r, err := c.render(); r != "" || err != nil {
		t.Error("Rendering vapor comment must returns empty string!")
	}

	// to native
	c = newComment("/* it'd be native")
	if c.commentType != C_TO_NATIVE {
		t.Error("It must be a 'to native' comment!")
	}

	if r, err := c.render(); r != "<!-- it'd be native -->\n" || err != nil {
		t.Error("To native comment rendering is broken!")
	}

	c.addContent("hold me too!")

	r, err := c.render()
	if err != nil ||
		removeMultipleSpaces(strings.Replace(r, "\n", "", -1)) != "<!-- it'd be nativehold me too! -->" {
		t.Error("Multiline native comment is broken!")
	}

	// native
	c = newComment("<!-- native html comment -->")
	if c.commentType != C_NATIVE {
		t.Error("It must be a native comment!")
	}

	if r, err := c.render(); r != "<!-- native html comment -->\n" || err != nil {
		t.Error("Native comment is broken!")
	}

	// variables
	name, value, err := parseVariable(removeComment("$my-var = test with // comments"))
	if err != nil {
		t.Error(err)
	}

	if name != "my-var" || value != "test with" {
		t.Error("Variable parsing with comment is broken!")
	}

	_, _, err = parseVariable(removeComment("$var = // it should be an error!"))
	if err == nil {
		t.Error("Variable parsing without value is not allowed!")
	}
}
