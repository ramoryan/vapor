// doctype
package vapor

type doctype struct {
	*element
}

func newDoctype(raw string) (*doctype, *vaporError) {
	e, err := newElement(raw)
	if err != nil {
		return nil, err
	}
	d := &doctype{element: e}
	d.name = "!DOCTYPE"
	d.addAttr("html", "")
	d.isVoid = true
	return d, nil
}
