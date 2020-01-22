package acronym

import (
	"bytes"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var acronym bytes.Buffer

	phrase := strings.Replace(s, "-", " ", -1)
	phrase = strings.Replace(phrase, "_", "", -1)
	ss := strings.Split(phrase, " ")

	for i := 0; i < len(ss); i++ {
		if len(ss[i]) != 0 {
			acronym.WriteByte(ss[i][0])
		}
	}
	return strings.ToUpper(acronym.String())
}
