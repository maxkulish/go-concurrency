package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	repeatFn := func(
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

	toInt := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}

	primeFinder := func(
		done <-chan interface{},
		numStream <-chan int,
	) <-chan interface{} {
		primeFinderStream := make(chan interface{})
		go func() {
			defer close(primeFinderStream)

			for n := range numStream {
				prime := true

				for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
					if n%i == 0 {
						prime = false
						break
					}
				}

				if prime {
					select {
					case <-done:
						return
					case primeFinderStream <- n:
					}
				}
			}
		}()
		return primeFinderStream
	}

	randNum := func() interface{} { return rand.Intn(50000000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, randNum))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))

}

func IsPrime(n int) bool {

	if n <= 1 {
		return false
	}

	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
