// vapor project vapor.go
package vapor

func ParseFile(fileName string) (string, error) {
	p := newParser()
	err := p.parseFile(fileName)

	if err != nil {
		return "", err
	}

	return p.output, nil
}
