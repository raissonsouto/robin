package src

func GetPrefixes(routines int) chan string {
	exponent := 1
	alphabetSize := alphabetLength

	for routines > alphabetSize {
		alphabetSize = alphabetSize * alphabetLength
		exponent++
	}

	return expandAlphabet(exponent)
}

func expandAlphabet(exponent int) chan string {

	lastTry, indexArray := setup(exponent)
	var combination string

	newValidChars := make(chan string)

	for lastTry != combination {
		rotate(indexArray)
		combination = getCombination(alphabet, indexArray)
		newValidChars <- combination
	}

	return newValidChars
}
