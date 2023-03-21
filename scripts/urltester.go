package scripts

import (
	"fmt"
	"io"
	"net/http"
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

	if resp.StatusCode == 200 { //200, 204, 301, 302, 400, 401, 403, 405, 500
		fmt.Printf("Found path: %s\n", path)
	} else if resp.StatusCode == 204 {

	} else if resp.StatusCode == 2 {

	} else if resp.StatusCode == 200 {

	}
}

func writeS() {

}
