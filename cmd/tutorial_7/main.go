package main

import (
	"fmt"
	"sync"
	"time"
)

// Basically it is a counter, we will add to the counter for each thread, once done
// We will decrement the counter. We then wait for the counter to reach 0 again
var waitGroup = sync.WaitGroup{}
// mutex lock when writing to our array so thay we donnt get corrupt data/race
// var mutex = sync.Mutex{}

// The other mutex locks other routines from being able to read from the results mutex
// We now also have read lock/unluck
var mutex = sync.RWMutex{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}  // slice

func dbCall(idx int) {
	// var delay float32 = rand.Float32() * 2000
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) *  time.Millisecond)
	fmt.Printf("\nThe result from the database is: %v", dbData[idx])

	save(dbData[idx])
	log()

	waitGroup.Done()
}

func save(result string) {
	// critical area
	// aquire a full lock, also if a read lock is currently aquired, we will wait (all locks must be clear)
	mutex.Lock()
	results = append(results, result)
	mutex.Unlock()
}

func log() {
	// check to see if there is currently a full lock (mutex.lock())
	// If yes, then wait
	// If no, we can get a read lock and proceed
	// mutliple go routines can hold read locks at the same time, these will only block a full lock
	mutex.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	mutex.RUnlock()
}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// use a go routine so we can execute concurrently
		// increment the counter in the wait group
		waitGroup.Add(1)
		// We spawn each concurrent thread
		go dbCall(i)
	}
	waitGroup.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}
