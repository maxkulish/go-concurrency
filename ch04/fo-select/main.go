package main

func main() {

	done := make(chan interface{})

	for {
		select {
		// if the done channel isn't closed, we'll exit the select statement
		case <- done:
			return
		default:
		// continue on the rest of our for loop's body
		}

		// Do some work in the loop
	}

	for {
		select {
		// if the done channel isn't close, execute the default clause instead
		case <- done:
			return
		default:
			// Do some work
		}
	}


}
