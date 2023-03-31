package src

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func testUrl(url string, path string) {
	fmt.Println(url + path)

	resp, err := http.Get(url + path)
	if err != nil {
		fmt.Printf("Error accessing URL: %s\n", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %s\n", err)
		}
	}(resp.Body)

	validStatusCodes := []int{200, 204, 301, 302, 401, 500}
	var log string

	if containsInt(resp.StatusCode, validStatusCodes) {
		log = logSuccessfulGuess(path, url, resp.StatusCode)
	}

	err = appendToFile(log, "")
	if err != nil {
		return
	}
}

func containsInt(intNum int, intList []int) bool {
	for _, validInt := range intList {
		if validInt == intNum {
			return true
		}
	}

	return false
}

func appendToFile(text string, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		printErrorFileNotFound()
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)

	if err != nil {
		return err
	}

	return nil
}
