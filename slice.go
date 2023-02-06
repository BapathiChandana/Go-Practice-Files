// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	//declare and initialize
	var a []int
	a = []int{10, 12, 2, 14, 6, 4}
	fmt.Println(a)
	println(&a)
	fmt.Println(a[0:2])
	fmt.Println(a[4])

	//declare & initialize in same line
	var b = []int{1, 2, 3, 4}
	fmt.Println(b)

	//initializing: short hand form
	c := []int{1, 2, 3, 4}
	fmt.Println(c)

	//assign 1 to another slice
	slice1 := []int{10, 20, 30, 40}
	slice2 := slice1
	fmt.Println(slice1)
	fmt.Println(slice2)

	//initialized all elements with zero value
	var d []int
	fmt.Println(d)

	//len and cap system functions
	fmt.Println(len(a), cap(a))

	//accessing individual element
	fmt.Println(b[3])

	//changing individual element
	slice1[2] = 50
	fmt.Println(slice1)

	//for loop:lenght hard cored
	for i := 0; i < 4; i++ {
		fmt.Print(c[i])
	}

}
