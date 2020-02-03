package twofer

import "fmt"

// ShareWith - `Two-fer` or `2-fer` is short for two for one. One for you and one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
