// main
package main

import (
	"fmt"

	"github.com/ramoryan/vapor"
)

func main() {
	out := vapor.ParseFile("./mytemplate.vapr")
	fmt.Println(out)
}
