package main

import (
	"fmt"
)

// ReURL custom url defination
type ReURL struct {
	Host   string `json:"host"`
	Port   string `json:":port"`
	Params map[string]string
}

func main() {
	fmt.Println("Nothing to see here")
}
