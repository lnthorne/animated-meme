package main

import "fmt"


func main() {
	var c = make(chan int)
	go process(c)
	// this process will wait until there is a value in the channel
	// fmt.Println(<-c)
	for i := range c {
		fmt.Println(i)
	}

	// WE can now store 5 ints in this channel at onece
	// var bufferChannel = make(chan int, 5)
}

func process(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}

	// closes the channel letting other processes know that we are done sending data
	close(c)
}
