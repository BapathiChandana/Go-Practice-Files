package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, b := range a {
		sum += b
	}
	fmt.Print(sum)
}
