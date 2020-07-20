package main

import (
	"fmt"
	"net/http"
)

// Get will perfrom a GET request on a given url
//
// Parmeters:
//	- `url` string : the url to perform the request at
//
// Returns:
//	- `*http.Response` - pointer to a `net/http` http.Response
//	- `error` - error
func Get(url string) (*http.Response, error) {
	// Perform the request and handle error
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error: could not perform get %s", err.Error())
	}

	// Return response
	return resp, nil
}

// Post will perfrom a POST request on a given url with given body
//
// Parmeters:
//	- `url` string : the url to perform the request at
//	- `body` ReBody : the body object created using `SendBody()`
//
// Returns:
//	- `*http.Response` - pointer to a `net/http` http.Response
//	- `error` - error
func Post(url string, body ReBody) (*http.Response, error) {
	// Perform the request and handle error
	resp, err := http.Post(url, body.ContentType, body.Content)

	if err != nil {
		return nil, fmt.Errorf("error: could not perform post %s", err.Error())
	}

	return resp, nil
}
