package raindrops

import "strconv"

// Covert - convert a number into a string that contains raindrop sounds corresponding to certain potential factors
func Convert(a int) string {

	retorno := ""

	if a == 1 {
		return "1"
	}

	if a%3 == 0 {
		retorno = "Pling"
	}

	if a%5 == 0 {
		retorno = retorno + "Plang"
	}

	if a%7 == 0 {
		retorno = retorno + "Plong"
	}

	if retorno == "" {
		retorno = strconv.Itoa(a)
	}
	return retorno
}
