// doctype
package vapor

type doctype struct {
	*element
}

func newDoctype(raw string) *doctype {
	d := &doctype{newElement(raw)}
	d.name = "!DOCTYPE"
	d.addAttr("html", "")
	d.isVoid = true
	return d
}
