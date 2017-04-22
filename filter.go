// filter
package vapor

import (
	"net/url"
	"strings"
)

func isFilter(s string) bool {
	return strings.HasPrefix(s, ":")
}

func resolveFilters(s string) string {
	if !isFilter(s) {
		return s
	}

	spc := strings.Index(s, " ")
	filters := s[:spc]
	content := s[spc+1:]

	filterNames := strings.Split(filters, ":")

	for _, name := range filterNames {
		content = resolveFilter(name, getVariableSafe(content))
	}

	return content
}

func resolveFilter(name, content string) string {
	if name == "upper" {
		return strings.ToUpper(content)
	} else if name == "lower" {
		return strings.ToLower(content)
	} else if name == "trim" {
		return strings.TrimSpace(content)
	} else if name == "title" {
		return strings.Title(content)
	} else if name == "capitalize" {
		return strings.ToUpper(string(content[0])) + content[1:]
	} else if name == "nl2br" {
		return strings.Replace(content, "\n", "<br>", -1)
	} else if name == "url_encode" {
		content, _ = url.QueryUnescape(content)

		return url.QueryEscape(content)
	}

	return content
}
