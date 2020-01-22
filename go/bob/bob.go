package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	if strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark && strings.HasSuffix(remark, "?") {
		return "Calm down, I know what I'm doing!"
	} else if strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark {
		return "Whoa, chill out!"
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	} else if remark == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}

}
