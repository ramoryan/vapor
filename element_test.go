package vapor

import (
	"testing"
)

func TestNewElement(t *testing.T) {
	r := "Element is broken!"

	// simple
	e, err := newElement("div")
	if err != nil || e.render() != "<div></div>\n" {
		t.Error(r)
	}

	// with one attr
	e, err = newElement(`div(id="my-id")`)
	if err != nil || e.render() != `<div id="my-id"></div>`+"\n" {
		t.Error(r)
	}

	// with multiple attrs
	e, err = newElement(`div(id="my-id" class="my-class")`)
	if err != nil || e.render() != `<div id="my-id" class="my-class"></div>`+"\n" {
		t.Error(r)
	}

	// --- shortcuts

	// #
	e, err = newElement("#my-id")
	if err != nil || e.render() != `<div id="my-id"></div>`+"\n" {
		t.Error(r)
	}

	// .
	e, err = newElement(".my-class")
	if err != nil || e.render() != `<div class="my-class"></div>`+"\n" {
		t.Error(r)
	}

	// #.
	e, err = newElement("#my-id.my-class")
	a := e.getAttributes()
	if err != nil || len(a) != 2 || !hasAttr(e, "id", "my-id") || !hasAttr(e, "class", "my-class") {
		t.Error(r)
	}

	// .#
	e, err = newElement(".my-class#my-id")
	a = e.getAttributes()
	if err != nil || len(a) != 2 || !hasAttr(e, "id", "my-id") || !hasAttr(e, "class", "my-class") {
		t.Error(r)
	}

	// shortcuts with tag
	e, err = newElement("input.my-inputClass#my-inputId")
	a = e.getAttributes()
	if err != nil || e.getName() != "input" || len(a) != 2 || !hasAttr(e, "id", "my-inputId") || !hasAttr(e, "class", "my-inputClass") {
		t.Error(r)
	}

	// shortcuts with other attributes
	e, err = newElement(`input.my-inputClass#my-inputId(style="border: 1px solid red;")`)
	a = e.getAttributes()
	if err != nil || e.getName() != "input" || len(a) != 3 ||
		!hasAttr(e, "id", "my-inputId") ||
		!hasAttr(e, "class", "my-inputClass") ||
		!hasAttr(e, "style", "border: 1px solid red;") {
		t.Error(r)
	}

	// cuts, attrs, variable interpolation, boolean attr
	clearStrStrMap(variables)
	setVariable("ghost", "that you can see")

	e, err = newElement(`input#my-input(class="#{ $ghost }" type="checkbox" checked)`)
	a = e.getAttributes()
	if err != nil || e.getName() != "input" || len(a) != 4 ||
		!hasAttr(e, "id", "my-input") ||
		!hasAttr(e, "class", "that you can see") ||
		!hasAttr(e, "type", "checkbox") ||
		!hasAttr(e, "checked", "") {
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
