// head
package vapor

type head struct {
	*element
}

func (h *head) addMeta(name, value string) *meta {
	m := newMeta("")
	m.addAttr(name, value)
	h.addChild(m)
	return m
}

func newHead(raw string) *head {
	h := &head{newElement(raw)}
	h.name = "head"

	h.addMeta("charset", "utf-8")
	h.addMeta("name", "viewport").addAttr("content", "width=device-width, initial-scale=1.0")

	return h
}
