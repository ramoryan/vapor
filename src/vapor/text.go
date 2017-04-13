// text
// https://pugjs.org/language/plain-text.html
package vapor

import (
	"strings"
)

type text struct {
	*element
}

func (t *text) render() string {
	return renderIndent(t.indent) + t.inlineText + "\n"
}

func newText(raw string) *text {
	t := &text{newElement(raw)}
	t.inlineText = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(t.raw), "|"))
	return t
}

func isText(s string) bool {
	if strings.HasPrefix(s, "|") {
		return true
	}

	return false
}
