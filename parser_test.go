package vapor

import (
	"testing"
)

func TestParser(t *testing.T) {
	p := newParser()
	err := p.parseFile("./invalidFile.vapr")
	if err == nil {
		t.Error("Parser must returns error!")
	}

	// multiline attributes
	s := []string{`
		input#my-id(
			type="checkbox"
			class="my-class"
			checked
			style="1px solid red"
			data-turboteddy="omg!"
			disabled
		)
	`}

	p = newParser()
	err = p.parseLines(s)
	if err != nil {
		t.Error(err)
	}

	if len(p.tree) != 1 {
		t.Error("Parser tree must contains one element!")
	}

	i := p.tree[0]

	if i.getName() != "input" ||
		!i.hasAttr("type", "checkbox") || !i.hasAttr("class", "my-class") ||
		!i.hasAttr("checked", "") || !i.hasAttr("style", "1px solid red") ||
		!i.hasAttr("data-turboteddy", "omg!") || !i.hasAttr("disabled", "") {
		t.Error("Multiline attribute parsing is broken!")
	}
}
