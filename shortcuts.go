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

			e.addAttrQ("href", value)
			e.addAttrQ("rel", "stylesheet")
			e.addAttrQ("type", "text/css")

			return e
		} else if name == "js" {
			e, _ := newElement("")
			e.name = "script"

			if !strings.HasSuffix(value, ".js") {
				value += ".js"
			}

			e.addAttrQ("src", value)

			return e
		} else if strings.HasPrefix(name, "og:") { // http://ogp.me/
			m, _ := newMeta("")
			m.addAttrQ("property", name)
			m.addAttrQ("content", value)

			return m
		} else if name == "keywords" || name == "author" || name == "description" {
			m, _ := newMeta("")
			m.addAttrQ("name", name)
			m.addAttrQ("content", value)

			return m
		}
	}

	return nil
}
