// vapor project vapor.go
package vapor

import (
	"errors"
	"fmt"
	"time"
)

// You can parse a vapor file and gets back the output string or a vaporError
func ParseFile(fileName string) (string, error) {
	start := time.Now()

	p := newParser()
	err := p.parseFile(fileName)

	if err != nil {
		return "", errors.New(err.String())
	}

	elapsed := time.Since(start)
	fmt.Printf("Running time: %s\n", elapsed)
	fmt.Println("Rendered vapor elem: " + intToStr(renderCount, ""))

	return p.output, nil
}

func AddStrVar(name, value string) *vaporError {
	// TODO: name and value len must be greater than 0!
	setVariable(name, value)
	return nil
}

func AddIntVar(name string, value int) *vaporError {
	// TODO: name len must greater than 0
	// value valid int
	setVariable(name, intToStr(value, ""))
	return nil
}

func AddStrSliceVar(name string, value []string) *vaporError {
	// TODO: is valid slice?
	s := make([]interface{}, len(value))
	for i, v := range value {
		s[i] = v
	}

	setVariable(name, s)
	return nil
}

func AddIntSliceVar(name string, value []int) *vaporError {
	s := make([]interface{}, len(value))
	for i, v := range value {
		s[i] = v
	}

	setVariable(name, s)
	return nil
}

func AddMapVar(name string, value map[string]interface{}) *vaporError {
	setVariable(name, value)
	return nil
}
