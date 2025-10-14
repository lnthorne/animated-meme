package main

import "fmt"

func main() {
	// pointers are initially nil
	var p *int32
	fmt.Println(p)

	var j *int32 = new(int32)

	// you can dereferecne the pointer as well
	fmt.Printf("The value that j points to is: %v", *j)
	*j = 20 // set the value j points to as 20


}
