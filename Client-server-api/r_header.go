package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://w3schools.com")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// for key, value := range r.Header {
// 			fmt.Println(key, ":", value)
// 		}
// 	})
// 	http.ListenAndServe(":8080", nil)
// }
// resp, err := http.Get("http://example.com/")
// if err != nil {
//     // handle error
// }
// defer resp.Body.Close()
// body, err := io.ReadAll(resp.Body)
