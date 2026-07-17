package handlershelp

import (
	"strings"
	"unicode"
)

// formatPersonName returns name with the first letter uppercased and the rest lowercased.
func FormatPersonName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return name
	}
	runes := []rune(strings.ToLower(name))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
