package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Custom-Header", "This is Go class")
		fmt.Fprintf(w, "Hello World")
	})

	http.ListenAndServe(":8080", nil)
}
