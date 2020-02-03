package reverse

//Reverse a string
//For example: input: "cool" output: "looc"
func Reverse(word string) string {
	reverse := ""
	if len(word) != 0 {
		wordSlice := []rune(word)
		for i := len(wordSlice) - 1; i >= 0; i-- {
			reverse += string(wordSlice[i])
		}
	}
	return reverse
}
