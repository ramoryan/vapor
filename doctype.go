// doctype
package vapor

type doctype struct {
	*element
}

func newDoctype(raw string) *doctype {
	e, _ := newElement(raw)
	d := &doctype{element: e}
	d.name = "!DOCTYPE"
	d.addAttr("html", "")
	d.isVoid = true
	return d
}
