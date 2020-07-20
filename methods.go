package main

import (
	"fmt"
	"net/http"
)

// Get will perfrom a GET request on a given url
func Get(url string) (*http.Response, error) {
	// Perform the request and handle error
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error: could not perform get %s", err.Error())
	}

	// Return response
	return resp, nil
}
