package vapor

import (
	"testing"
)

func TestNewAttribute(t *testing.T) {
	var a attribute
	var err *vaporError

	_, err = newAttribute("id", `"my-id"`)
	if err != nil {
		t.Error(err.errMsg)
	}

	_, err = newAttribute("", "")
	if err == nil {
		t.Error("Empty name is not allowed!")
	}

	_, err = newAttribute("id", "my-id")
	if err == nil {
		t.Error("Quotes must be used!")
	}

	setVariable("variable", "value") // make dummy
	_, err = newAttribute("id", "$variable")
	if err != nil {
		t.Error("Variable without quotes must be allowed!" + err.String())
	}

	a, err = newAttribute("id", "")
	if a.name != "id" {
		t.Error("Name without value must be allowed!")
	}
}

func TestRender(t *testing.T) {
	a, _ := newAttribute("id", `"my-id"`)
	if a.render() != `id="my-id"` {
		t.Error("Attribute rendering with value has been broken!")
	}

	a, _ = newAttribute("checked", "")
	if a.render() != "checked" {
		t.Error("Single name attribute rendering has been broken!")
	}
}

func TestParseAttribute(t *testing.T) {
	e := "Parse attribute has been broken!"

	name, value := parseAttribute("checked")
	if name != "checked" || len(value) != 0 {
		t.Error(e)
	}

	name, value = parseAttribute(`class="my-class`)
	if name != "class" && value != "my-class" {
		t.Error(e)
	}
}
