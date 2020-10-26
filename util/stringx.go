package util

import (
	"strings"
)

func UnTitle(s string) string {
	if len(s) == 0 {
		return s
	}
	first := s[:1]
	ret := strings.ToLower(first) + s[1:]
	return ret
}
