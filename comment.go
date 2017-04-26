// comment
package vapor

import (
	"reflect"
	"strings"
)

const (
	C_VAPOR     = 1
	C_TO_NATIVE = 2
	C_NATIVE    = 3
)

type comment struct {
	*element
	content     []string
	commentType int
}

func (c *comment) render() string {
	if c.commentType == C_VAPOR {
		return ""
	}

	spc := renderIndent(c.indent + 8)
	s := renderIndent(c.indent)

	if c.commentType != C_NATIVE {
		s += "<!-- "
	}

	for i, line := range c.content {
		if i > 0 {
			s += spc
		}

		s += line

		if i < len(c.content)-1 {
			s += "\n"
		}
	}

	if c.commentType != C_NATIVE {
		s += " -->"
	}

	return s + "\n"
}

func getCommentType(s string) int {
	// saját comment, nincs kimenete
	if strings.HasPrefix(s, "//") {
		return C_VAPOR
	} else if strings.HasPrefix(s, "/*") { // html commentté alakítjuk
		return C_TO_NATIVE
	} else if strings.HasPrefix(s, "<!") { // natív html comment
		return C_NATIVE
	}

	return -1
}

func (c *comment) addContent(s string) {
	if c.commentType != C_VAPOR {
		c.content = append(c.content, strings.TrimSpace(s))
	}
}

func isComment(s string) bool {
	return getCommentType(s) > -1
}

func isCommentType(v vaporizer) bool {
	t := reflect.TypeOf(v).String()

	if t == "*vapor.comment" {
		return true
	}

	return false
}

func newComment(raw string) *comment {
	c := &comment{element: &element{raw: raw}} // ne newElement-et használj!
	c.setIndent(calcIndent(raw))
	c.commentType = getCommentType(strings.TrimSpace(raw))

	if c.commentType != C_VAPOR {
		content := strings.TrimSpace(strings.Replace(raw, "/*", "", -1))

		if len(content) > 0 {
			c.addContent(content)
		}
	}

	return c
}

func removeComment(s string) string {
	tmp := strings.TrimSpace(s)
	if pos := strings.Index(tmp, "//"); pos > 0 {
		return s[:pos]
	}

	return s
}
