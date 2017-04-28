// utils
package vapor

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// var isAllowedStr = regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString
var rgxMultipleSpaces = regexp.MustCompile(`[\s\p{Zs}]{2,}`)

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

func log(f interface{}) {
	fmt.Println(f)
}

func typeof(f interface{}) reflect.Kind {
	return reflect.TypeOf(f).Kind()
}

func isSlice(f interface{}) bool {
	return typeof(f) == reflect.Slice
}

func isStr(f interface{}) bool {
	return typeof(f) == reflect.String
}

func isMap(f interface{}) bool {
	return typeof(f) == reflect.Map
}

func isInt(f interface{}) bool {
	return typeof(f) == reflect.Int
}

func isIterateable(f interface{}) bool {
	if !isSlice(f) && !isStr(f) && !isMap(f) {
		return false
	}

	return true
}

// https://gobyexample.com/collection-functions
func sliceMap(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func splitAndTrim(s, sep string) []string {
	split := strings.Split(s, sep)
	if len(split) > 0 {
		split = sliceMap(split, strings.TrimSpace)
	}
	return split
}

// Helpers for testing

func removeMultipleSpaces(s string) string {
	return rgxMultipleSpaces.ReplaceAllString(s, "")
}

func clearStrStrMap(m map[string]string) {
	for k := range m {
		delete(m, k)
	}
}
