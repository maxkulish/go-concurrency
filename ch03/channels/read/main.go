package main

import (
	"fmt"
	"sync"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)

	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	fmt.Printf("(%v): %v\n", ok, integer)

	intStream1 := make(chan int)
	go func() {
		defer close(intStream1)
		for i := 1; i <= 5; i++ {
			intStream1 <- i
		}
	}()

	for integer := range intStream1 {
		fmt.Printf("%v ", integer)
	}

	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// the goroutine waits until it is told it can continue
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	// unblocking all the goroutines simultaneously
	close(begin)
	wg.Wait()
}
