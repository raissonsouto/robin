package scripts

import (
	"sync"
)

const url = "https://go.dev/"

var wg sync.WaitGroup

const qtdRoutines = 50
const pathSize = 10

func main() {

	workBalancer := workBalancer(validChars, qtdRoutines)

	for i := 0; i < qtdRoutines; i++ {
		initialString := <-workBalancer
		go fuzz(url, initialString, pathSize, validChars)
	}
}
