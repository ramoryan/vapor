package vapor

import (
	"testing"
)

func TestNewHead(t *testing.T) {
	h, err := newHead("")
	if err == nil {
		c := h.getChildren()

		if len(c) != 2 ||
			c[0].getName() != "meta" || c[1].getName() != "meta" ||
			!c[0].hasAttr("charset", "utf-8") ||
			!c[1].hasAttr("name", "viewport") ||
			!c[1].hasAttr("content", "width=device-width, initial-scale=1.0") {
			t.Error("Head has been broken!")
		}
	} else {
		t.Error(err)
	}
}

func TestHeadAddChild(t *testing.T) {
	h, _ := newHead("")

	for _, tag := range validHeadChildren {
		e, err := newElement(tag)
		if err != nil {
			t.Error(err)
		}

		err = h.addChild(e)
		if err != nil {
			t.Error(err)
		}
	}

	e, err := newElement("div")
	err = h.addChild(e)
	if err == nil {
		t.Error("Not allowed tag in head!")
	}
}
