package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memCosumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	// this goroutine won't exit until the process is finished
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memCosumed()
	for i:= numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memCosumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
