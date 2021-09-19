package main

import "fmt"

func main() {
	// ad hoc confinement: data slice is available from both
	// the loopData function and the loop over the handleData channel
	// by convention we're only accessing it from the loopData function
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}