// meta
package vapor

type meta struct {
	*element
}

func newMeta(raw string) *meta {
	e, _ := newElement(raw)
	m := &meta{element: e}
	m.name = "meta"
	m.isVoid = true
	return m
}
