package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~:/[]@!$&'()*+,;="
const size = 20
const routines = 100
const url = "https://go.dev/"

var wg sync.WaitGroup

func prodPossibilities(buffer chan string, path string) {
	defer wg.Done()

	if len(path) > size {
		return
	}

	buffer <- path

	for _, char := range validChars {
		prodPossibilities(buffer, path+string(char))
	}
}

func testPossibilities(buffer chan string) {
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
	buffer := make(chan string)
	i := 0

	fmt.Println("[*] Starting Robin")

	go prodPossibilities(buffer, "")

	for _, char := range validChars {
		wg.Add(1)
		go prodPossibilities(buffer, string(char))
	}

	for i < routines {
		wg.Add(1)
		go testPossibilities(buffer)
		i++
	}

	wg.Wait()
	fmt.Println("[*] Finished")
}
