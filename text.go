// text
// https://pugjs.org/language/plain-text.html
package vapor

import (
	"strings"
)

type text struct {
	*element
}

func (t *text) render() (string, *vaporError) {
	return renderIndent(t.indent) + t.inlineText + "\n", nil
}

func newText(raw string) (*text, *vaporError) {
	e, _ := newElement(raw)
	t := &text{element: e}
	txt, err := interpolateVariables(strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(t.raw), "|")))
	if err != nil {
		return nil, err
	}

	t.inlineText = txt
	return t, nil
}

func isText(s string) bool {
	return strings.HasPrefix(s, "|")
}
