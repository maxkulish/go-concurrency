package main

import "fmt"

func main()  {
	stringStream := make(chan string)
	go func() {
		// example of deadlock
		// fatal error: all goroutines are asleep - deadlock!
		if 0 != 1 {
			return
		}
		stringStream <- "Hello channels!"
	}()
	fmt.Println(<-stringStream)
}
