package main

import (
	"errors"
	"fmt"
)

func main() {
	const printValue string = "Hello World"
	printMe(printValue)

	var result int32 = division(4, 2)
	fmt.Println(result)

	var value, remainder = multipleReturnDivision(4, 3)
	fmt.Printf("The result of the integer divison is %v with remainder of %v\n", value, remainder)

	// Check if error was returned 
	var value2, remainder2, err = divisionWithError(3, 0)
	if (err != nil) {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The result of the integer division is %v with a remainder of %v", value2, remainder2)
	}
}

func printMe(value string) {
	fmt.Println(value)
}

func division(numerator int32, denominator int32) int32 {
	return numerator / denominator
}

func multipleReturnDivision(numerator int32, denominator int32) (int32, int32) {
	var result int32 = numerator / denominator
	var remainder int32 = numerator % denominator

	return result, remainder
}

func divisionWithError(numerator int32, denominator int32) (int32, int32, error) {
	// Default value of nil
	var err error
	if (denominator == 0) {
		err = errors.New("cannot divide by 0")
		return 0, 0, err 
	}
	var result int32 = numerator / denominator
	var remainder int32 = numerator % denominator

	return result, remainder, err
}