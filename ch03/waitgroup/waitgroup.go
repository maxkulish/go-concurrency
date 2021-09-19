package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main()  {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // will block the main goroutine until all goroutines have indicated they have exited
	fmt.Println("All goroutines complete.")

	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("hello from %v!\n", id)
	}

	const numGreeters = 5
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
