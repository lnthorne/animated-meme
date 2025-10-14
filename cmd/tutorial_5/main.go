package main

import "fmt"

type owner struct {
	name string
}
// creating a struct
type gasEngine struct {
	mpg uint8
	gallons uint8
	ownerInfo owner
}

// This is now a method for the gasEnging struct
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}


// The interface now allows us to pass either engine into the canMakeIt func as both structs fulfill the interface
type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if (miles <= e.milesLeft()) {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You will not make it")
	}
}


func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons:15, ownerInfo: owner{"Alex"}}
	myEngine.gallons = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)

	// anonymous struct, non reusable
	var myEngine2 = struct {
		mpg uint8
		gallons uint8
	}{24, 26}

	fmt.Println(myEngine2.gallons, myEngine2.mpg)
	

	canMakeIt(myEngine, 50)
}
