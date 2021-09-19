package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		// instantiate the channel within the lexical scope of the chanOwner cunction
		// it confines the write aspect of this channel to prevent other goroutines from writing to it
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// receive a read-only copy of at int channel
	// confine usage of the channel within the consume function to only reads
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	// receive the read aspect of the channel
	// we're able to pass it into the consumer
	results := chanOwner()
	consumer(results)
}
