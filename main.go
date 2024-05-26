package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server...")
	httpPort := ":8080"
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
