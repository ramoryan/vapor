// vapor project vapor.go
package vapor

func ParseFile(fileName string) string {
	p := newParser()
	p.parseFile(fileName)
	return p.output
}
