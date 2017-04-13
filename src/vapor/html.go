// html
package vapor

import (
	"strings"
)

type html struct {
	*element
}

func newHtml(raw string) *html {
	h := &html{newElement(raw)}
	h.name = "html"

	if strings.Index(raw, "lang") < 0 {
		h.addAttr("lang", "en")
	}

	return h
}
