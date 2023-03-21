package scripts

import "fmt"

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numeric = "0123456789"
const special = "-._~:/[]@!$&'()*+,;="

const validChars = lower + upper + numeric + special

func printInitialMessage() {
	fmt.Println("")
	fmt.Println("STARTING ROBIN FUZZING")
	fmt.Println("")
	fmt.Println("url:")
	fmt.Println("routines quantity:")
	fmt.Println("guess max size:")
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
