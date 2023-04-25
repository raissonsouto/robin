package src

func AlphabetGenerator() {
	defer Wg.Done()
	exponent := 1
	alphabetSize := alphabetLength

	for QtdRoutines > alphabetSize {
		alphabetSize = alphabetSize * alphabetLength
		exponent++
	}

	lastTry, indexArray := setup(exponent)
	var combination string

	for lastTry != combination {
		rotate(indexArray)
		combination = getCombination(indexArray)
		prefixChan <- combination
	}

	deltaGuessSize = maxGuessSize - exponent
}
