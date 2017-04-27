// main
package main

import (
	"fmt"

	"github.com/ramoryan/vapor"
)

func main() {
	vapor.AddStrVar("vaporStr", "karamel")
	vapor.AddIntVar("vaporInt", 956)

	/*strSlice := []string{"a", "b", "c"}
	vapor.AddStrSliceVar("vaporSlice", strSlice)*/

	intSlice := []int{956, 1845, 1956}
	vapor.AddIntSliceVar("vaporSlice", intSlice)

	out, err := vapor.ParseFile("./mytemplate.vapr")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
}
