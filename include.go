// include
package vapor

import (
	"io/ioutil"
	"strings"
)

type htmlContent struct {
	*element
}

func (h *htmlContent) render() (res string) {
	indent := renderIndent(h.indent)
	lines := strings.Split(h.inlineText, "\n")

	for _, s := range lines {
		res += indent + s + "\n"
	}

	return res
}

func newHtmlContent(s string) *htmlContent {
	h := &htmlContent{element: &element{}}
	h.inlineText = s
	return h
}

func isInclude(str string) bool {
	return strings.HasPrefix(str, "include")
}

func include(str string) *htmlContent {
	path := strings.TrimSpace(strings.TrimPrefix(str, "include"))

	if strings.HasSuffix(path, ".html") {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			// hibakezelés
		}

		h := newHtmlContent(string(b))
		return h
	} else if strings.HasSuffix(path, ".vapr") {
		p := newParser()
		p.parseFile(path)
		h := newHtmlContent(p.output)
		return h
	} else {
		// ERROR
		return nil
	}
}
