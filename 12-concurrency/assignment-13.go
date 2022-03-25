package main

import (
	"fmt"
	"time"
)

func main() {
	//print the fibonacci series until the user hits ENTER key
	doneCh := make(chan bool)
	dataCh := genFibonacci(doneCh)
	go func() {
		var input string
		fmt.Scanln(&input)
		doneCh <- true
	}()
	for no := range dataCh {
		fmt.Println(no)
	}

}

/* Keep generating the fibonacci series */
func genFibonacci(done chan bool) <-chan int {
	dataCh := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-done:
				fmt.Println("done")
				close(dataCh)
				break LOOP
			case dataCh <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
	}()
	return dataCh
}
