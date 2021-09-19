package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultSream := chanOwner()
	for result := range resultSream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving")
}
