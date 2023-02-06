package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/hello2", handler2)
	http.ListenAndServe("Localhost:8082", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Chandana"))
}
func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Chandana bapathi"))
}
