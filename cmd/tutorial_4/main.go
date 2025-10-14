package main

import "fmt"

func main() {
	// fixed arrays of length 3
	var intArr [3]int32
	fmt.Println(intArr[0])
	intArr[1] = 123
	fmt.Println(intArr[1:3]) // gets elements in index 1 and 2

	// reference the memeory address 
	fmt.Println(&intArr[1])
}
