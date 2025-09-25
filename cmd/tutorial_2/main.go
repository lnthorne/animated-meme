package main

import "fmt"

func main() {
	// int8, int16, int32, int64
	// int will default to 32 or 64 bits depending on your architecture 
	// also have unsigned ints: uint8, uint16, uint32, uint64
	var intNum int = 0
	intNum = intNum + 1
	fmt.Println(intNum)

	// float32
	var floatNum float64 = 123.456
	fmt.Println(floatNum)


	var myString string = "Hello \n World"
	fmt.Println(myString)

	var myString2 string = "Hello" + " " + "World"
	fmt.Println(myString2)

	var myRune rune = 'A'
	fmt.Println(myRune)

	// Default values
	// All ints and floats default to 0
	var intNum2 int
	fmt.Println(intNum2) // default 0

	var myRune2 rune
	fmt.Println(myRune2) // default 0

	// String default is empty string ""
	var myString3 string
	fmt.Println(myString3) // ""

	// booleans the default is false 
	var myBoolean bool
	fmt.Println(myBoolean)


	// Declare variables
	// explicit 
	var myType string = "foo"

	// Implicit
	var myType2 = "hello"

	// Shorthand 
	myType3 := "World"

	fmt.Println(myType, myType2, myType3)


	// init multiple vars at the same time 
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// constants 
	// init the same as using var, just cannot change its value once created. It also needs to be initialized
	const myConst string = "const value"

}