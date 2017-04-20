// shortcuts
package vapor

import (
	"strings"
)

func resolveShortcut(s string) vaporizer {
	if pos := strings.Index(s, " "); pos > 0 {
		name := s[:pos]
		value := s[pos+1:]

		if name == "css" {
			e := newElement("")
			e.name = "link"
			e.addAttr("href", value)
			e.addAttr("rel", "stylesheet")
			e.addAttr("type", "text/css")

			return e
		} else if name == "js" {
			e := newElement("")
			e.name = "script"
			e.addAttr("src", value)

			return e
		} else if strings.HasPrefix(name, "og:") { // http://ogp.me/
			m := newMeta("")
			m.addAttr("property", name)
			m.addAttr("content", value)

			return m
		}
	}

	return nil
}
