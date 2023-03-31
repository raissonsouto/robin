package src

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numeric = "0123456789"
const special = "-._~:/[]@!$&'()*+,;="

const ValidChars = lower + upper + numeric + special

func printInitialMessage(url string, routines int, s int) {
	fmt.Println("")
	fmt.Println("STARTING ROBIN FUZZING")
	fmt.Println("")
	fmt.Println("url: " + url)
	fmt.Println("routines quantity: " + string(rune(routines)))
	fmt.Println("guess max size: " + string(rune(s)))
}

func printErrorUrlNotSpecified() {
	fmt.Println("Error: URL parameter is required")
}

func printErrorURLNotFound() {
	fmt.Println("Error: The URL were not found.")
}

func printErrorMaxSizeToSmall() {
	fmt.Println("Error: The maxSize should be >= 1.")
}

func printErrorFileNotFound() {
	fmt.Println("Error: The path or file name specified not found or created.")
}

func logSuccessfulGuess(path string, url string, httpResponse int) string {
	row := string(rune(httpResponse)) + "," + url + "," + path + "\n"
	return row
}
