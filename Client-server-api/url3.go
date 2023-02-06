package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("http://www.gmail.com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", *u)
}
