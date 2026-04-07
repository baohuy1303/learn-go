package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{} // wait group so that the program doesnt just stops
var mutex = sync.RWMutex{} // Mutex lock
var dbData = []string{"id1", "id2", "id3", "id4"}
var result = []string{}

func main() {

	var t0 = time.Now()

	for i:=0; i < len(dbData); i++{
		waitGroup.Add(1) // add each call to the wait group
		go dbCall(i) // goroutine call
	}
	waitGroup.Wait() // wait for all groups to finish
	fmt.Println("Total execution time of non-readlock: ", time.Since(t0), "\n\n")

	var t1 = time.Now()
	result = []string{}
	for i:=0; i < len(dbData); i++{
		waitGroup.Add(1) // add each call to the wait group
		go dbCallReadLocks(i) // goroutine call
	}
	waitGroup.Wait() // wait for all groups to finish
	fmt.Println("Total execution time of readlock: ", time.Since(t1))
}

func dbCall(i int){
	time.Sleep(time.Duration(500) * time.Millisecond)

	// Only 1 write at a time, while still printing sequentially
	mutex.Lock()
	fmt.Println("The result from the database is: " + dbData[i])
	result = append(result, dbData[i])
	fmt.Println("The current result is: ", result) // + operator for strings only works when both operands are strings
	mutex.Unlock()

	waitGroup.Done()
}

func dbCallReadLocks(i int){
	time.Sleep(time.Duration(500) * time.Millisecond)

	fmt.Println("The result from the database is: " + dbData[i])

	//Lock the result so none can write
	mutex.Lock()
	result = append(result, dbData[i])
	mutex.Unlock()

	// if there exists a lock (any lock), then it won't print
	// This ensures that the result is done writing, then it starts printing
	// Prevents weird printing if result is still being written, and ensure sequential
	mutex.RLock()
	fmt.Println("The current result is: ", result)
	mutex.RUnlock()

	waitGroup.Done()
}

// Goroutines can occupy all the cores in your CPU, if a task is light
// it releases the core early, allowing others to take over, so it could theoretically
// run thousand times more instances than the number of cores you have.
// BUT, if the task is computational or heavy, it takes time to release a CPU core
// Others have to wait, making the time increase linearly depending on your num of coures.