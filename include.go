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

func include(str string) (*htmlContent, error) {
	path := strings.TrimSpace(strings.TrimPrefix(str, "include"))

	if strings.HasSuffix(path, ".html") {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, newVaporError("Cannot include " + path)
		}

		h := newHtmlContent(string(b))
		return h, nil
	} else {
		if !strings.HasSuffix(path, ".vapr") {
			path += ".vapr"
		}

		p := newParser()
		err := p.parseFile(path)

		if err != nil {
			return nil, err
		}

		h := newHtmlContent(p.output)
		return h, nil
	}

	return nil, nil
}
