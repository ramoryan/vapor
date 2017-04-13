// main
package main

import (
	"fmt"
	"vapor"
)

func main() {
	out := vapor.ParseFile("./mytemplate.vapr")
	fmt.Println(out)
}
