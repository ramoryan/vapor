// vaporizer
package vapor

type vaporizer interface {
	render() (string, *vaporError)
	getIndent() int
	getParent() vaporizer
	setParent(v vaporizer)
	addChild(v vaporizer) *vaporError
	getAttributes() attrMap
	getChildren() []vaporizer
	getName() string
	setIndent(indent int)
	addAttr(name, value string) *vaporError
	needMultilineAttrs() bool
	closeMultilineAttrs()
	hasAttr(attrName, attrValue string) bool
	appendTree(t vaporTree)
	reIndent()
	parse() (vaporTree, *vaporError)
}
