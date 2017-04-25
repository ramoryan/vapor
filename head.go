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

func (h *head) addMeta(name, value string) (*meta, *vaporError) {
	m, _ := newMeta("")

	err := m.addAttr(name, value)
	if err != nil {
		return nil, err
	}

	err = h.addChild(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (h *head) addChild(v vaporizer) *vaporError {
	name := v.getName()

	sort.Strings(validHeadChildren)
	i := sort.SearchStrings(validHeadChildren, name)

	if i < len(validHeadChildren) && validHeadChildren[i] == name {
		h.element.addChild(v)
	} else {
		return newVaporError(ERR_HEAD, 1, "Not valid tag: "+name)
	}

	return nil
}

func newHead(raw string) (*head, *vaporError) {
	e, err := newElement(raw)
	if err != nil {
		return nil, err
	}

	h := &head{element: e}
	h.name = "head"

	_, err = h.addMeta("charset", `"utf-8"`)
	if err != nil {
		return nil, err
	}

	var m *meta
	m, err = h.addMeta("name", `"viewport"`)
	if err != nil {
		return nil, err
	}

	m.addAttr("content", `"width=device-width, initial-scale=1.0"`)

	return h, nil
}
