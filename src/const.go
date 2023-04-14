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

var alphabet = lower + upper + numeric + special
var alphabetLength int

func setAlphabet(regex string) {
	switch regex {
	case "a":
		alphabet = lower
	case "A":
		alphabet = upper
	case "0":
		alphabet = numeric
	case "$":
		alphabet = special
	case "Aa", "aA":
		alphabet = lower + upper
	case "A0", "0A":
		alphabet = upper + numeric
	case "A$", "$A":
		alphabet = upper + special
	case "a0", "0a":
		alphabet = lower + numeric
	case "a$", "$a":
		alphabet = lower + special
	case "0$", "$0":
		alphabet = numeric + special
	case "Aa0", "A0a", "aA0", "a0A", "0Aa", "0aA":
		alphabet = lower + upper + numeric
	default:
		alphabet = lower + upper + numeric + special
	}
	alphabetLength = len(alphabet)
}

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
