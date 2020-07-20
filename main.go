package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// SendBody passes a body to the url used for POST, DELETE etc. Only JSON is supported for now
//
// Parameters:
//	- `t` string : the body type to pass into request only JSON is supported for now
//	- `bod` interface{} : the body content to pass into request
func SendBody(t string, bod interface{}) {
	// Check for supported types
	switch strings.ToLower(t) {
	case "json":
		{
			// TODO: Send JSON as body
			fmt.Println("Type is json!")
			break
		}
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
		return nil, fmt.Errorf("error: JSON decoder error:\n%s", err.Error())
	}

	return converted, nil
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return
}
