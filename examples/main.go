// main
package main

import (
	"fmt"

	"github.com/ramoryan/vapor"
)

func main() {
	out, err := vapor.ParseFile("./mytemplate.vapr")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
}
