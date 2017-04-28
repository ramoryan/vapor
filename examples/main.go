// main
package main

import (
	"fmt"

	"github.com/ramoryan/vapor"
)

func main() {
	vapor.AddStrVar("vaporStr", "karamel")
	vapor.AddIntVar("vaporInt", 956)

	str := "abcdefgh"
	vapor.AddStrVar("vaporMainStr", str)

	strSlice := []string{"aaa", "bbb", "ccc"}
	vapor.AddStrSliceVar("vaporStrSlice", strSlice)

	intSlice := []int{956, 1848, 1956}
	vapor.AddIntSliceVar("vaporIntSlice", intSlice)

	out, err := vapor.ParseFile("./mytemplate.vapr")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
}
