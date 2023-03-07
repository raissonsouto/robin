package main

/*

package main

import (
	"fmt"
	"net/http"
	"sync"
)

const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~:/[]@!$&'()*+,;="

func prodPossibilities() {

}

func testPossibilities() {}

func findPath(url string, path string, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(path) > 2 {
		return
	}

	resp, err := http.Get(url + path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println(path)
	}

	for _, char := range validChars {
		wg.Add(1)
		go findPath(url, path+string(char), wg)
	}
}

func main() {
	// variaveis
	//smooth := false

	//routines = 10

	fmt.Println("Starting Robin ...")
	url := "https://overthewire.org/"

	var wg sync.WaitGroup
	wg.Add(1)
	go findPath(url, "", &wg)

	wg.Wait()
	fmt.Println("All goroutines have finished")
}

*/
