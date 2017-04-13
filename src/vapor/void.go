package vapor

import (
	"strings"
)

type void struct {
	*element
}

var voidElements = []string{
	"area",
	"base",
	"basefont",
	"bgsound",
	"br",
	"col",
	"command",
	"embed",
	"frame",
	"hr",
	"image",
	"img",
	"input",
	"isindex",
	"keygen",
	"link",
	"menuitem",
	"meta",
	"nextid",
	"param",
	"source",
	"track",
	"wbr",
}

func newVoidElement(raw string) *void {
	v := &void{newElement(raw)}
	v.isVoid = true
	return v
}

func isVoidElement(s string) bool {
	for _, name := range voidElements {
		if strings.HasPrefix(s, name) {
			return true
		}
	}

	return false
}
