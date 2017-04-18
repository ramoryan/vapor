// meta
package vapor

type meta struct {
	*element
}

func newMeta(raw string) *meta {
	m := &meta{newElement(raw)}
	m.name = "meta"
	m.isVoid = true
	return m
}
