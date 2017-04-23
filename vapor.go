// vapor project vapor.go
package vapor

import (
	"errors"
)

func ParseFile(fileName string) (string, error) {
	p := newParser()
	err := p.parseFile(fileName)

	if err != nil {
		return "", errors.New(err.String())
	}

	return p.output, nil
}
