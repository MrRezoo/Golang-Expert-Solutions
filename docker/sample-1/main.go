// main.go
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, Docker!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
