package src

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func testUrl(path string) {
	fmt.Println("d" + targetUrl + path)

	resp, err := http.Get(targetUrl + path)
	if err != nil {
		fmt.Printf("Error accessing URL: %s\n", err)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %s\n", err)
		}
	}(resp.Body)

	if checkStatusCode(validStatusCodes, resp.StatusCode) {
		log := logSuccessfulGuess(path, targetUrl, resp.StatusCode)
		fmt.Println(log)
	}

	err = writeToLogFile("teste")
	if err != nil {
		return
	}
}

func checkStatusCode(validStatusList []int, status int) bool {
	for _, num := range validStatusList {
		if num == status {
			return true
		}
	}
	return false
}

func writeToLogFile(content string) error {

	var err error

	if _, err = os.Stat(outputFile); os.IsNotExist(err) {
		if _, err = os.Create(outputFile); err != nil {
			return fmt.Errorf("failed to create results.txt: %v", err)
		}
	}

	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open results.txt: %v", err)
	}
	defer file.Close()

	if _, err = fmt.Fprintln(file, content); err != nil {
		return fmt.Errorf("failed to write to results.txt: %v", err)
	}

	return nil
}
