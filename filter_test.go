package vapor

import (
	"testing"
)

func TestResolveFilters(t *testing.T) {
	e := "Filter has been broken!"

	s, err := resolveFilters("random string")
	if s != "random string" || err != nil {
		t.Error(e)
	}

	s, err = resolveFilters(":notvalid text")
	if err == nil {
		t.Error("Not valid filter did not make error!")
	}

	s, _ = resolveFilters(":upper text")
	if s != "TEXT" {
		t.Error("Upper filter broken!")
	}

	s, _ = resolveFilters(":lower TEXT")
	if s != "text" {
		t.Error("Lower filter broken!")
	}

	s, _ = resolveFilters(":trim      text      ")
	if s != "text" {
		t.Error("Trim filter broken!")
	}

	s, _ = resolveFilters(":title text text")
	if s != "Text Text" {
		t.Error("Title filter broken!")
	}

	s, _ = resolveFilters(":capitalize text text")
	if s != "Text text" {
		t.Error("Capitalize filter broken!")
	}

	s, _ = resolveFilters(":nl2br text\n")
	if s != "text<br>" {
		t.Error("nl2br filter broken!")
	}

	s, _ = resolveFilters(":url_encode http://www.golang.com")
	if s != "http%3A%2F%2Fwww.golang.com" {
		t.Error("url_encode filter broken!")
	}

	s, _ = resolveFilters(":url_encode http%3A%2F%2Fwww.golang.com")
	if s != "http%3A%2F%2Fwww.golang.com" {
		t.Error("url_encode filter broken!")
	}

	// multiple filters
	s, _ = resolveFilters(":trim:upper      text     ")
	if s != "TEXT" {
		t.Error(s)
	}
}
