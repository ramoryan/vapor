// vaporizer
package vapor

type vaporizer interface {
	render() string
	getIndent() int
	getParent() vaporizer
	setParent(v vaporizer)
	addChild(v vaporizer)
	setIndent(indent int)
	addAttr(name, value string)
}
