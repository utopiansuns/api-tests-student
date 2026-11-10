package utils

import "strings"

// MixCase чередует регистр букв: "горный" → "ГоРнЫй"
func MixCase(s string) string {
	var b strings.Builder
	i := 0
	for _, r := range s {
		ch := string(r)
		if i%2 == 0 {
			b.WriteString(strings.ToUpper(ch))
		} else {
			b.WriteString(strings.ToLower(ch))
		}
		i++
	}
	return b.String()
}
