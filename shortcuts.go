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
			e, _ := newVoidElement("")
			e.name = "link"

			if !strings.HasSuffix(value, ".css") {
				value += ".css"
			}

			e.addAttr("href", value)
			e.addAttr("rel", "stylesheet")
			e.addAttr("type", "text/css")

			return e
		} else if name == "js" {
			e, _ := newElement("")
			e.name = "script"

			if !strings.HasSuffix(value, ".js") {
				value += ".js"
			}

			e.addAttr("src", value)

			return e
		} else if strings.HasPrefix(name, "og:") { // http://ogp.me/
			m, _ := newMeta("")
			m.addAttr("property", name)
			m.addAttr("content", value)

			return m
		} else if name == "keywords" || name == "author" || name == "description" {
			m, _ := newMeta("")
			m.addAttr("name", name)
			m.addAttr("content", value)

			if name == "description" && len(value) > 160 {
				// error
			}

			return m
		}
	}

	return nil
}
