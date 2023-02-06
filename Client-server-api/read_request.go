package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	Name    string
	Address string
	Id      string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	c := student{}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%#v", c)
}
