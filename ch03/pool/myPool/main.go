package main

import (
	"fmt"
	"sync"
)

func main()  {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return struct {}{}
		},
	}

	// invoke the New function
	myPool.Get()
	instance := myPool.Get()
	// Put instance back in the pool. This increases the available number of instances to one
	myPool.Put(instance)
	// Reuse the instance previously allocated and put it back in the pool.
	// The New function will not be invoked
	myPool.Get()
}