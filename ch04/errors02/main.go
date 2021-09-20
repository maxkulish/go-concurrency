package main

import (
	"fmt"
	"net/http"
)

func main() {

	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{
					Error:    err,
					Response: resp,
				}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	errCount := 0
	urls := []string{"https://www.google.com", "https://badhost", "https://memo.ua", "a", "b", "c"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			errCount++
			fmt.Printf("error: %v\n", result.Error)
			fmt.Println("==>>")
			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
		fmt.Printf("Headers: %v\n", result.Response.Header)
		fmt.Println("==>>")
	}
}
