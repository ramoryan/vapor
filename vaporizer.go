// vaporizer
package vapor

type vaporizer interface {
	render() string
	getIndent() int
	getParent() vaporizer
	setParent(v vaporizer)
	addChild(v vaporizer) *vaporError
	getAttributes() []attribute
	getChildren() []vaporizer
	getName() string
	setIndent(indent int)
	addAttr(name, value string) *vaporError
	needMultilineAttrs() bool
	closeMultilineAttrs()
}
