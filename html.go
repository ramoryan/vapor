// html
package vapor

import (
	"strings"
)

type html struct {
	*element
}

func newHtml(raw string) (*html, *vaporError) {
	e, err := newElement(raw)
	if err != nil {
		return nil, err
	}

	h := &html{element: e}
	h.name = "html"

	if strings.Index(raw, "lang") < 0 {
		err := h.addAttr("lang", `"en"`)
		if err != nil {
			return nil, err
		}
	}

	return h, nil
}
