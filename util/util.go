package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ByteRequest returns a byte slice from a url
func ByteRequest(u string) ([]byte, error) {
	res, err := http.Get(u)
	if err != nil {
		fmt.Printf("Error: Cannot resolve URL '%s': %s", u, err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: Unable to read body of request: %s", err)
		return nil, err
	}
	return body, nil
}
