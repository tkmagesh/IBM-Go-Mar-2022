/* alternative implementation for demo-12.go */
package main

import (
	"fmt"
)

//Share Memory for communication

func main() {

	fmt.Println("main started")

	ch := add(100, 200)
	result := <-ch

	fmt.Println(result)
	fmt.Println("main completed")
}

func add(x, y int) chan int {
	ch := make(chan int)
	result := x + y
	ch <- result
	return ch
}
