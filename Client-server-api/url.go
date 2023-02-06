package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL)
}

func main() {
	http.HandleFunc("w3schools.com", handler)
	http.ListenAndServe(":8085", nil)
}
