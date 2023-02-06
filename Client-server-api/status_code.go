package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err == nil {
		http.Error(w, "error message", http.StatusInternalServerError)
		////w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(b))

	w.Write([]byte("hello"))

}
func handleRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe("localhost:4205", nil))
}
func main() {
	handleRequest()
}
