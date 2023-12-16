package utility

import (
	"html"
	"regexp"
)

var textareaRegex = regexp.MustCompile(`(?s)<textarea.*?>(.*?)</textarea>`)

func FindSrcFromMediawiki(s string) string {
	r := textareaRegex.FindStringSubmatch(s)
	if r == nil || len(r) != 2 {
		return ""
	}
	rr := html.UnescapeString(r[1])
	return rr
}
