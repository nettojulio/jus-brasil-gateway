package client

import (
	"fmt"
	"io"
	"net/http"
)

func GetRawHTML(url string) ([]byte, error) {
	fmt.Println("Making request to external URL")
	resp, _ := http.Get(url)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when making URL request")
		return []byte(nil), err
	}
	return body, nil
}
