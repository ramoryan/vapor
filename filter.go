// filter
package vapor

import (
	"net/url"
	"strings"
)

func isFilter(s string) bool {
	return strings.HasPrefix(s, ":")
}

func resolveFilters(s string) (string, *vaporError) {
	if !isFilter(s) {
		return s, nil
	}

	spc := strings.Index(s, " ")
	filters := s[:spc]
	content := s[spc+1:]

	filterNames := strings.Split(filters, ":")

	for _, name := range filterNames {
		if len(name) <= 0 {
			continue // TODO: hack?
		}

		res, err := resolveFilter(name, getVariableSafe(content).(string))
		if err != nil {
			return "", err
		}

		content = res
	}

	return content, nil
}

func resolveFilter(name, content string) (string, *vaporError) {
	if name == "upper" {
		return strings.ToUpper(content), nil
	} else if name == "lower" {
		return strings.ToLower(content), nil
	} else if name == "trim" {
		return strings.TrimSpace(content), nil
	} else if name == "title" {
		return strings.Title(content), nil
	} else if name == "capitalize" {
		return strings.ToUpper(string(content[0])) + content[1:], nil
	} else if name == "nl2br" {
		return strings.Replace(content, "\n", "<br>", -1), nil
	} else if name == "url_encode" {
		content, _ = url.QueryUnescape(content)

		return url.QueryEscape(content), nil
	}

	return "", newVaporError(ERR_FILTER, 1, "Filter is not defined: "+name)
}
