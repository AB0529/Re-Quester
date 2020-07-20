package quester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// ReHeader the header structure that get passed in
type ReHeader struct {
	Key   string
	Value string
}

// ReBody the body type which will be sent empty interface to force user to use `SendBody` function
type ReBody struct {
	ContentType string
	Content     io.Reader
	Headers     []ReHeader
}

// SendBody passes a body to the url used for POST, DELETE etc. Only JSON is supported for now
//
// Parameters:
//	- `t` string : the body type to pass into request only JSON is supported for now
//	- `bod` interface{} : the body content to pass into request
//	- `headers []ReHeader` : the headers that will be passed into the request
func SendBody(t string, bod interface{}, headers []ReHeader) ReBody {
	// Check for supported types
	switch strings.ToLower(t) {
	case "json":
		{
			requestBody, err := json.Marshal(bod)
			if err != nil {
				panic(err)
			}

			// Send nil headers
			if headers == nil {
				return ReBody{ContentType: "application/json", Content: bytes.NewBuffer(requestBody), Headers: nil}
			}

			// Send request with headers
			return ReBody{ContentType: "application/json", Content: bytes.NewBuffer(requestBody), Headers: headers}
		}
	default:
		panic("error: could not convert to any types")
	}
}

// BodyToMap will convert a net/http response body into a map
//
// Parameters:
//	- `body` io.ReadCloser : the response body that will be converted
//
// Returns:
//	- map[string]interface{} : The result of the body into a map
func BodyToMap(body io.ReadCloser) (map[string]interface{}, error) {
	var converted map[string]interface{}
	data, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(data, &converted)

	if err != nil {
		return nil, fmt.Errorf("error: JSON decoder %s", err.Error())
	}

	defer body.Close()
	return converted, nil
}

// Client the http client which are defined when running methods
var Client *http.Client

// Get will perfrom a GET request on a given url
//
// Parmeters:
//	- `url` string : the url to perform the request at
//
// Returns:
//	- `*http.Response` - pointer to a `net/http` http.Response
//	- `error` - error
func Get(url string, body ReBody) (*http.Response, error) {
	// Create the request
	Client := &http.Client{}

	// Check if body is passed in
	if body.ContentType != "" {
		req, err := http.NewRequest("GET", url, body.Content)

		if err != nil {
			return nil, fmt.Errorf("error: could not perform get %s", err.Error())
		}

		// Set the headers if they exist
		if body.Headers != nil {
			for _, content := range body.Headers {
				req.Header.Set(content.Key, content.Value)
			}
		}

		// Perform request with headers
		resp, err := Client.Do(req)

		return resp, nil
	}

	// Perform the request without headers
	req, err := http.NewRequest("GET", url, body.Content)
	resp, err := Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error: could not perform get %s", err.Error())
	}

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
	// Create the request
	Client := &http.Client{}
	req, err := http.NewRequest("POST", url, body.Content)

	// Set the headers if they exist
	if body.Headers != nil {
		for _, content := range body.Headers {
			req.Header.Set(content.Key, content.Value)
		}
	}

	// Perform the request
	resp, err := Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error: could not perform post %s", err.Error())
	}

	return resp, nil
}
