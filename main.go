package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// ReBody the body type which will be sent empty interface to force user to use `SendBody` function
type ReBody interface{}

// SendBody passes a body to the url used for POST, DELETE etc. Only JSON is supported for now
//
// Parameters:
//	- `t` string : the body type to pass into request only JSON is supported for now
//	- `bod` interface{} : the body content to pass into request
func SendBody(t string, bod interface{}) (ReBody, error) {
	// Check for supported types
	switch strings.ToLower(t) {
	case "json":
		{
			// TODO: Send JSON as body
			fmt.Println("Type is json!")
			break
		}
	}

	return nil, fmt.Errorf("error: could not convert to any types")
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

func main() {
	res, _ := Get("https://jsonplaceholder.typicode.com/todos/1")
	defer res.Body.Close()

	// Convert response to map
	data, err := BodyToMap(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(data["id"]) // 1
}
