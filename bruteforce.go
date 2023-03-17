package main

import "fmt"

func startClock(initialString string, maxSize int, alphabet string) {

	diffSize := maxSize - len(initialString)
	alphabetLength := len(alphabet) - 1

	var token string
	var lastTry string
	var indexArray = make([]int, diffSize)

	for i := 0; i < diffSize; i++ {
		indexArray[i] = -1
		lastTry += string(alphabet[alphabetLength])
	}

	fmt.Println(lastTry)

	for lastTry != token {
		rotate(indexArray, alphabetLength)
		token = getToken(alphabet, indexArray)
		fmt.Println(initialString + token)
	}
}

func getToken(alphabet string, indexArray []int) string {
	var result string

	for _, index := range indexArray {
		if index == -1 {
			continue
		}

		result += string(alphabet[index])
	}

	return result
}

func rotate(indexArray []int, alphabetLength int) {
	for index := 1; index < len(indexArray)-1; index++ {

		if indexArray[index] == alphabetLength {
			indexArray[index] = 0
			indexArray[index+1]++
		}
	}
	if indexArray[0] == alphabetLength {
		indexArray[0] = 0
		indexArray[0+1]++
	} else {
		indexArray[0]++
	}
}

func main() {
	alphabet := "abcd"
	startClock("/", 8, alphabet)
}
