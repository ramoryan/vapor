// for
package vapor

import (
	//"strconv"
	"strings"
)

// for 1 to 4

func isLoop(s string) bool {
	if strings.HasPrefix(s, "for") && strings.Index(s, " to ") > 0 {
		return true
	}

	return false
}

func handleLoop(s string, indent int) vaporizer {
	c := newContainer(indent)
	return c
}
