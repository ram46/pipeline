package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func updateUser() string {
	return "Simple Helloworld server is running on port 8080"
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println(updateUser())
	http.ListenAndServe(":8080", nil)
}
