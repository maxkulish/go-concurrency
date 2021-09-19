package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
		) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <- strings:
					// Do something interesting
					fmt.Println(s)
				// If done channel has been signaled, we return from the goroutine
				case <- done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		// Cancel the operation after 1 second
		time.Sleep(5*time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	// join the goroutine spawned from doWork with the main goroutine
	<-terminated
	fmt.Println("Done")
}