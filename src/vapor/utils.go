// utils
package vapor

import (
	"strconv"
	"strings"
)

func renderIndent(i int) string {
	return strings.Repeat(" ", i)
}

func calcIndent(s string) int {
	i := 0

	for _, c := range s {
		if c == 32 {
			i += 1
		} else if c == 9 {
			i = ((i + 8) / 8) * 8
		} else if c != 13 {
			break
		}
	}

	return i
}

func unquote(s string) string {
	un, err := strconv.Unquote(s)

	if err != nil {
		return s
	}

	return un
}
