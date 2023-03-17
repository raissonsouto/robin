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

func insertInitialStringsInQueue() {
	defer wg.Done()

	for _, char := range validChars {
		validCharsChannel <- char
	}

	fmt.Println("[+] Valid characters inserted in the queue")
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
*/
