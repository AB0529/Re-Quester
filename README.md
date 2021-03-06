# Re-Quester

This is a simple wrapper for Go's **net/http** libary created with the purpose of having a simpler way to make and send requests.

# Functions

- `Get(url string, body ReBody)` - Will send a GET request to the url provided

    - returns **(\*http.Response, error)**

- `Post(url string, body ReBody)` - Will send a POST request with the provided body

    - returns **(\*http.Response, error)**
    
- `SendBody(type string, body interface{}, headers []ReHeader)` - Creates body to send as type

    - returns **ReBody**

- `BodyToMap(body io.ReadCloser)` - Converts a **net/http** response body into a map

    - returns **(map[string]interface{}, error)**

# Types
- `ReBody` - the body type which will be sent, can be built with **SendBody()**
    - `ContentType string` - The content type header which will be sent i.e **application/json** (only json supported for now)

    - `Content io.Reader` - The request body whicih will be sent

    - `Headers []ReHeader` - The headers which will be sent to the request

- `ReHeader` - the header that will be passed
    - `Key string` - the header key
    - `Value string` - the header value

```go
// The data that will be used as a body for requests
dataToSend := map[string]string{
    "title":  "Hello",
    "body":   "Body",
    "userId": "123", }

// Construct body using SendBody, can be made manually as well!
myBodyToSend := SendBody("json", dataToSend, nil)
// ...
```

```go
// The data that will be used as a body for requests
dataToSend := map[string]string{
    "title":  "Hello",
    "body":   "Body",
    "userId": "123", }
headersToSend := []ReHeader{ReHeader{"Cool-Key", "Epic-Value"}}

// Construct body using SendBody, can be made manually as well!
myBodyToSend := SendBody("json", dataToSend, headersToSend)
// ...
```

# Examples

## Making Requests

```go
package main

import (
    "fmt"
    "github.com/AB0529/Re-Quester"
)

func main() {
    res, _ := Get("https://jsonplaceholder.typicode.com/todos/1", ReBody{})
    defer res.Body.Close()

    // Convert response to map
    data, _ := BodyToMap(res.Body, nil)

    fmt.Println(data["id"]) // 1
}
```

## Sending Requests

```go
package main

import (
    "fmt"
    "github.com/AB0529/Re-Quester"
)

func main() {
    // Doesn't have to be a map!
    dataToSend := map[string]string{"title":  "Hello", "body":   "Body", "userId": "123", }
    headersToSend := []ReHeader{ReHeader{"Cool-Key", "Epic-Value"}}

    // Perform request
    res, _ := Post("https://jsonplaceholder.typicode.com/posts", SendBody("json", dataToSend, headersToSend))
    defer res.Body.Close()

    // Convert to map
    data, _ := BodyToMap(res.Body)

    fmt.Println(data["title"]) // Hello
}
```