// main
package main

import (
	"fmt"

	"github.com/ramoryan/vapor"
)

func main() {
	vapor.AddStrVar("vaporStr", "karamel")
	vapor.AddIntVar("vaporInt", 956)

	strSlice := []string{"a", "b", "c"}
	vapor.AddStrSliceVar("vaporSlice", strSlice)
	out, err := vapor.ParseFile("./mytemplate.vapr")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
}
