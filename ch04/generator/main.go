package main

import (
	"fmt"
	"math/rand"
)

func main() {
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	// This pipeline stage will only take the first num items off of its incoming valueStream and then exit
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	repeatFn :=
		func(
			done <-chan interface{},
			fn func() interface{},
		) <-chan interface{} {
			valueStream := make(chan interface{})
			go func() {
				defer close(valueStream)
				for {
					select {
					case <-done:
						return
					case valueStream <- fn():
					}
				}
			}()
			return valueStream
		}

	toString := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1, 2, 3), 10) {
		fmt.Printf("%v\n", num)
	}

	randNum := func() interface{} { return rand.Int() }

	for n := range take(done, repeatFn(done, randNum), 10) {
		fmt.Println(n)
	}

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}

	fmt.Printf("message: %s...", message)
}
