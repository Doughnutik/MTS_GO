package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// data := `{"name": "Artem", "age": 20}`
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.)
	fmt.Println(string(body))
}
