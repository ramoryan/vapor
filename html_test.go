package vapor

import (
	"testing"
)

func TestNewHtml(t *testing.T) {
	e := "HTML rendering has been broken! "

	h, err := newHtml("html")
	if h.render() != `<html lang="en"></html>`+"\n" {
		if err != nil {
			t.Error(err)
		} else {
			t.Error(e + h.render())
		}
	}

	h, err = newHtml(`html(lang="hu")`)
	if h.render() != `<html lang="hu"></html>`+"\n" {
		if err != nil {
			t.Error(err)
		} else {
			t.Error(e + h.render())
		}
	}
}
