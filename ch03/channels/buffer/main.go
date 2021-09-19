package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// it's a little faster than writing to stdout directly
	var stdoutBuff bytes.Buffer
	// ensure that the buffer is written out to stdout before the process exits
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}

}