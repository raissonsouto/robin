package src

import (
	"fmt"
	"strings"
)

func Fuzz() {
	defer Wg.Done()

	var combination string
	lastTry, indexArray := setup(deltaGuessSize)

	teste := <-prefixChan
	fmt.Println(teste)

	for lastTry != combination {
		rotate(indexArray)
		combination = getCombination(indexArray)
		fmt.Println(combination)
		//testUrl(combination)
	}
}

func getCombination(indexArray []int) string {
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
