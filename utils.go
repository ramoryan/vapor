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
	if un, err := strconv.Unquote(s); err == nil {
		return un
	}

	return s
}

func isBetweenQuotes(s string) bool {
	return (strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`)) || (strings.HasPrefix(s, `'`) && strings.HasSuffix(s, `'`))
}

func isLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func strToInt(s string, def int) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}

	return def
}

func intToStr(i int, def string) string {
	return strconv.Itoa(i)
	/*if s, err := strconv.Itoa(i); err == nil {
		return i
	}

	return def*/
}
