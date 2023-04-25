package src

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

var targetUrl string
var QtdRoutines int
var deltaGuessSize int
var maxGuessSize int
var smooth, detach, commonPathsOnly, help bool

var prefixChan = make(chan string, 1)

const outputFile = "paths.txt"

var validStatusCodes = []int{200, 204, 301, 302, 401, 500}

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

func PrintInitialMessage() {
	fmt.Println("\nSTARTING ROBIN FUZZING")
	fmt.Println("\n| url: " + targetUrl)
	fmt.Printf("| routines quantity: %v\n", QtdRoutines)
	fmt.Printf("| guess max size: %v\n", maxGuessSize)
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
