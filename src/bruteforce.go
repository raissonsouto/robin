package src

import (
	"strings"
)

func Fuzz(url string, exponent int) {

	var combination string
	lastTry, indexArray := setup(exponent)

	for lastTry != combination {
		rotate(indexArray)
		combination = getCombination(alphabet, indexArray)
		testUrl(url, combination)
	}

	defer Wg.Done()
}

func loop(exponent int) { // funcs ...interface{}

	var combination string
	lastTry, indexArray := setup(exponent)

	for lastTry != combination {
		rotate(indexArray)
		combination = getCombination(alphabet, indexArray)
	}
}

func getCombination(alphabet string, indexArray []int) string {
	var result strings.Builder

	for _, index := range indexArray {
		if index == -1 {
			continue
		}

		result.WriteByte(alphabet[index])
	}

	return result.String()
}

func rotate(indexArray []int) {

	indexArray[0]++

	for index := 0; index < len(indexArray)-1; index++ {

		if indexArray[index] == alphabetLength {
			indexArray[index] = 0
			indexArray[index+1]++
		}
	}
}

func setup(exponent int) (string, []int) {
	var lastTry string
	var indexArray = make([]int, exponent)

	for i := 0; i < exponent; i++ {
		indexArray[i] = -1
		lastTry += string(alphabet[alphabetLength-1])
	}

	return lastTry, indexArray
}
