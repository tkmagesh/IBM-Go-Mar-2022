/* demo-08.go without WaitGroup */
package main

import (
	"fmt"
)

var counter = 0 //race

func main() {

	fmt.Println("main started")
	doneCh := make(chan bool)
	for i := 0; i < 100; i++ {
		go fn(doneCh)
	}
	for i := 0; i < 100; i++ {
		<-doneCh
	}

	fmt.Println("counter = ", counter)
	fmt.Println("main completed")
}

func fn(done chan bool) {
	counter = counter + 1
	fmt.Println(counter)
	done <- true
	fmt.Println("fn completed")
}
