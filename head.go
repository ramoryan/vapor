// head
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head
package vapor

import (
	"sort"
)

type head struct {
	*element
}

var validHeadChildren = []string{"title", "base", "link", "style", "meta", "script", "noscript"}

func (h *head) addMeta(name, value string) *meta {
	m := newMeta("")
	m.addAttr(name, value)
	h.addChild(m)
	return m
}

func (h *head) addChild(v vaporizer) {
	name := v.getName()

	sort.Strings(validHeadChildren)
	i := sort.SearchStrings(validHeadChildren, name)

	if i < len(validHeadChildren) && validHeadChildren[i] == name {
		h.element.addChild(v)
	} else {
		// éktelen haragra gerjedés :)
	}
}

func newHead(raw string) *head {
	h := &head{newElement(raw)}
	h.name = "head"

	h.addMeta("charset", "utf-8")
	h.addMeta("name", "viewport").addAttr("content", "width=device-width, initial-scale=1.0")

	return h
}
