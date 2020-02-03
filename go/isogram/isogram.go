package isogram

import (
	"strings"
)

// IsIsogram - Determine if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	if len(word) != 0 {
		word = removeCharacterEspecials(word)
		for _, l := range word {
			if strings.Count(word, string(l)) != 1 {
				return false
			}
		}
	}
	return true
}

func removeCharacterEspecials(word string) string {
	word = strings.ToUpper(word)
	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, " ", "", -1)
	return word
}
