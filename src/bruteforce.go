package src

import (
	"strings"
)

func Fuzz(url string, initialString string, maxSize int, alphabet string) {

	exponent := maxSize - len(initialString)
	alphabetLength := len(alphabet)

	var token string
	var lastTry string
	var indexArray = make([]int, exponent)

	for i := 0; i < exponent; i++ {
		indexArray[i] = -1
		lastTry += string(alphabet[alphabetLength-1])
	}

	for lastTry != token {
		rotate(indexArray, alphabetLength)
		token = getCombination(alphabet, indexArray)
		testUrl(url, token)
	}

	defer Wg.Done()
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

func rotate(indexArray []int, alphabetLength int) {

	indexArray[0]++

	for index := 0; index < len(indexArray)-1; index++ {

		if indexArray[index] == alphabetLength {
			indexArray[index] = 0
			indexArray[index+1]++
		}
	}
}
