package vapor

import (
	"testing"
)

func TestNewHtml(t *testing.T) {
	e := "HTML rendering has been broken! "

	h, err := newHtml("html")
	if err != nil {
		t.Error(err)
	}

	if r, err := h.render(); r != `<html lang="en"></html>`+"\n" || err != nil {
		if err != nil {
			t.Error(err)
		} else {
			t.Error(e + r)
		}
	}

	h, err = newHtml(`html(lang="hu")`)
	if err != nil {
		t.Error(err)
	}

	if r, err := h.render(); r != `<html lang="hu"></html>`+"\n" || err != nil {
		if err != nil {
			t.Error(err)
		} else {
			t.Error(e + r)
		}
	}
}
