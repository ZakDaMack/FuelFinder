package slug

import (
	"regexp"
	"strings"
)

var invalidChars = regexp.MustCompile("[^%sa-zA-Z0-9]")

func Slugify(s string) string {
	s = invalidChars.ReplaceAllString(s, "-")
	s = strings.Trim(s, "")
	s = strings.Trim(s, "-")
	s = strings.ToLower(s)
	return s
}
