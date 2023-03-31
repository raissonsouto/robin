package src

import (
	"math"
)

func WorkBalancer(alphabet string, routines int) chan string {

	exponent := 1
	alphabetSize := len(alphabet)

	for routines > alphabetSize {
		alphabetSize = int(math.Pow(float64(alphabetSize), 2))
		exponent++
	}

	return expandValidChars(alphabet, exponent)
}

func expandValidChars(alphabet string, exponent int) chan string {

	alphabetLength := len(alphabet)
	newValidChars := make(chan string)

	var indexArray = make([]int, exponent)

	for i := 0; i < exponent; i++ {
		indexArray[i] = -1
	}

	combinations := int(math.Pow(float64(alphabetLength), float64(exponent)))

	for i := 0; combinations > i; i++ {
		rotate(indexArray, alphabetLength)
		newValidChars <- getCombination(alphabet, indexArray)
	}

	return newValidChars
}
