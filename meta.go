// meta
package vapor

type meta struct {
	*element
}

func newMeta(raw string) (*meta, *vaporError) {
	e, err := newElement(raw)
	if err != nil {
		return nil, err
	}
	m := &meta{element: e}
	m.name = "meta"
	m.isVoid = true
	return m, nil
}
