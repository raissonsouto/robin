package main

/*
import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

const pathSize = 10
const qtdProdRoutines = 50
const qtdConsRoutines = 50
const url = "https://go.dev/"
const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numeric = "0123456789"
const special = "-._~:/[]@!$&'()*+,;="
const validChars = lower + upper + numeric + special

var validCharsChannel = make(chan int32)
var generatedPaths = make(chan string)
var wg sync.WaitGroup

func insertValidCharsInQueue() {
	defer wg.Done()

	for _, char := range validChars {
		validCharsChannel <- char
	}

	fmt.Println("[+] Valid characters inserted in the queue")
}

func prodPossibilities() {
	defer wg.Done()

	for {
		letter := <-validCharsChannel

		rootPath := string(letter)
		generatedPaths <- rootPath

		for _, char := range validChars {
			path := rootPath + string(char)
			if len(path) > pathSize {
				break
			}
			generatedPaths <- path
		}
	}
}

func prodPossibility(rootPath string) {
	for _, char := range validChars {
		if len(path) > pathSize {
			break
		}
		prodPossibilities(path)
	}
}

func testPossibilities() {
	defer wg.Done()

	for {
		possibility := <-buffer
		fmt.Println(url + possibility)
		resp, err := http.Get(url + possibility)
		if err != nil {
			fmt.Printf("Error accessing URL: %s\n", err)
			continue
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Printf("Error closing response body: %s\n", err)
			}
		}(resp.Body)

		if resp.StatusCode == 200 {
			fmt.Printf("Found path: %s\n", possibility)
			break
		}
	}
}

func main() {
	fmt.Println("[*] Starting Robin")
	fmt.Printf("[*] Using %d producer routines and %d consumer routines\n", qtdProdRoutines, qtdConsRoutines)

	wg.Add(1)
	go insertValidCharsInQueue()

	for i := 0; i < qtdProdRoutines; i++ {
		wg.Add(1)
		go prodPossibilities()
	}

	for i := 0; i < qtdConsRoutines; i++ {
		wg.Add(1)
		go testPossibilities()
	}

	wg.Wait()
	fmt.Println("[*] Finished")
}

func brute(maxLength int) {



	for i := 0; i < maxLength; i++ {
		for _, char := range validChars {
			for
		}
	}
}
*/
