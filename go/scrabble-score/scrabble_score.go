package scrabble

import "strings"

// Score - Given a word, compute the Scrabble score for that word.
func Score(word string) int {
	var result = 0
	word = strings.ToUpper(word)
	for _, n := range word {
		switch n {
		case 'D', 'G':
			result += 2
		case 'B', 'C', 'M', 'P':
			result += 3
		case 'F', 'H', 'V', 'W', 'Y':
			result += 4
		case 'K':
			result += 5
		case 'J', 'X':
			result += 8
		case 'Q', 'Z':
			result += 10
		default:
			result++
		}
	}
	return result
}
