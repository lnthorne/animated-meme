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

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	// Maps will always return an default value of the value type even if the key does not esist

	// Maps can also return an optional boolean so you can check if the key exists and the value is correct
	var age, ok = myMap2["Adam"]
	if (ok) {
		fmt.Println(age)
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for idx, value := range intArr {
		fmt.Printf("Index: %v, value: %v", idx, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	delete(myMap2, "Adam")
}
