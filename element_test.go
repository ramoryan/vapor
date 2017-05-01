package vapor

import (
	"testing"
)

func BenchmarkNewElement(b *testing.B) {
	clearVariables()
	setVariable("ghost", "that you can see")

	for i := 0; i < b.N; i++ {
		_, err := newElement(`div#my-div(class="#{ $ghost }" disabled style="1px solid black") Ez meg itt szöveg`)
		if err != nil {
			break
		}
	}
}

func TestNewElement(t *testing.T) {
	r := "Element is broken!"

	// simple
	e, err := newElement("div")
	if err != nil {
		t.Error(r)
	}
	if s, err := e.render(); s != "<div></div>\n" || err != nil {
		t.Error(r)
	}

	// with one attr
	e, err = newElement(`div(id="my-id")`)
	if err != nil {
		t.Error(r)
	}

	if ren, err := e.render(); ren != `<div id="my-id"></div>`+"\n" || err != nil {
		t.Error(r)
	}

	// with multiple attrs
	e, err = newElement(`div(id="my-id" class="my-class")`)
	if err != nil {
		t.Error(err)
	}

	if ren, err := e.render(); ren != `<div class="my-class" id="my-id"></div>`+"\n" || err != nil {
		t.Error(r)
	}

	// --- shortcuts

	// #
	e, err = newElement("#my-id")
	if err != nil {
		t.Error(r)
	}

	if ren, err := e.render(); ren != `<div id="my-id"></div>`+"\n" || err != nil {
		t.Error(r)
	}

	// .
	e, err = newElement(".my-class")
	if err != nil {
		t.Error(r)
	}

	if ren, err := e.render(); ren != `<div class="my-class"></div>`+"\n" || err != nil {
		t.Error(r)
	}

	// #.
	e, err = newElement("#my-id.my-class")
	a := e.getAttributes()
	if err != nil || len(a) != 2 || !e.hasAttr("id", "my-id") || !e.hasAttr("class", "my-class") {
		t.Error(r)
	}

	// .#
	e, err = newElement(".my-class#my-id")
	a = e.getAttributes()
	if err != nil || len(a) != 2 || !e.hasAttr("id", "my-id") || !e.hasAttr("class", "my-class") {
		t.Error(r)
	}

	// shortcuts with tag
	e, err = newElement("input.my-inputClass#my-inputId")
	a = e.getAttributes()
	if err != nil || e.getName() != "input" || len(a) != 2 || !e.hasAttr("id", "my-inputId") || !e.hasAttr("class", "my-inputClass") {
		t.Error(r)
	}

	// shortcuts with other attributes
	e, err = newElement(`input.my-inputClass#my-inputId(style="border: 1px solid red;")`)
	a = e.getAttributes()
	if err != nil || e.getName() != "input" || len(a) != 3 ||
		!e.hasAttr("id", "my-inputId") ||
		!e.hasAttr("class", "my-inputClass") ||
		!e.hasAttr("style", "border: 1px solid red;") {
		t.Error(r)
	}

	// cuts, attrs, variable interpolation, boolean attr
	clearVariables()
	setVariable("ghost", "that you can see")

	e, err = newElement(`input#my-input(class="#{ $ghost }" type="checkbox" checked)`)
	a = e.getAttributes()
	if err != nil || e.getName() != "input" || len(a) != 4 ||
		!e.hasAttr("id", "my-input") ||
		!e.hasAttr("class", "that you can see") ||
		!e.hasAttr("type", "checkbox") ||
		!e.hasAttr("checked", "") {
		t.Error(r)
	}

	// attributes with same name error
	sameNameError := "Attributes with same name is not allowed!"

	e, err = newElement(`div(id="my-id" id="my-id")`)
	if err != nil {
		t.Error(sameNameError)
	}

	e, err = newElement(`#my-id(id="my-id")`)
	if err != nil {
		t.Error(sameNameError)
	}

	e, err = newElement(`.my-class(class="my-class")`)
	if err != nil {
		t.Error(sameNameError)
	}

	setVariable("attrName", "id")
	e, err = newElement(`#my-id(#{$attrName}="my-id")`)
	if err != nil {
		t.Error(sameNameError)
	}

	// --- try to break it!

	e, err = newElement(`div(alma=)`)
	if err == nil {
		t.Error("Attribute equation without value is not allowed!")
	}

	// TODO: error!
	// e, err = newElement(`div(alma=""a)`)

	// TODO: error!
	// e, err = newElement(`div(alma=""`)

	// TODO: error!
	// e, err = newElement(`div()`)

	e, err = newElement("#")
	if err == nil {
		t.Error("Id shortcut without value is not allowed!")
	}

	e, err = newElement(".")
	if err == nil {
		t.Error("Class shortcut without value is not allowed!")
	}
}
