package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// create our condition using a standard sync.Mutex
	c := sync.NewCond(&sync.Mutex{})
	// create a slice with a length of zero
	// we know we'll eventually and 10 items, so we instantiate it with a capacity of 10
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		// enter the critical section for the condition, so we can modify data
		c.L.Lock()
		// simulate dequeuing an item by reassigning the head of the slice to the second item
		queue = queue[1:]
		fmt.Println("Removed from queue")
		// exit the condition's critical section
		c.L.Unlock()
		// we let a goroutine waiting on the condition know that something has occured
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			// suspend the main goroutine until a signal on the condition has been sent
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		// create a new goroutine that will dequeue an element after one second
		go removeFromQueue(1*time.Second)
		c.L.Unlock()
	}
}
