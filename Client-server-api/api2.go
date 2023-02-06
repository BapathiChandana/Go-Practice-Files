package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	w.Write([]byte("Hello Chandana"))

}
func handleRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe("localhost:4200", nil))
}
func main() {
	handleRequest()
}
