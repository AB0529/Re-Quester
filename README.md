# Re-Quester

This is a simple wrapper for Go's `net/http` libary created with the purpose of having a simpler way to make and send requests.

# Functions

- `Get(url string)` - Will send a GET request to the url provided
- `Post(url string, body ReBody)` - Will send a POST request with the provided body
- `SendBody(type string, body interface{})` - Constructs a body to send as type provided
- `BodyToMap(body io.ReadCloser)` - Converts a `net/http` response body into a map

# Types
- `ReBody` - the body type which will be sent

```go
type ReBody interface{}
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
	res, _ := Get("https://jsonplaceholder.typicode.com/todos/1")
	defer res.Body.Close()

	// Convert response to map
	data, _ := BodyToMap(res.Body)

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
    // Doesn't have to be map!
    DataToPost := make(map[string][interface{}]) 
    DataToPost = append(DataToPost["name"], "Bob")

    res, _ := Post("https://jsonplaceholder.typicode.com/posts", SendBody("json", DataToPost))
    // ...
}
```