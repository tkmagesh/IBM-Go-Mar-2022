/* select construct */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "routine-1 done"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "routine-2 done"
	}()

	/*
		d1 := <-ch1
		fmt.Println("operation - 1 completed", d1)
		d2 := <-ch2
		fmt.Println("operation - 2 completed", d2)
	*/
	for i := 0; i < 2; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println("operation - 1 completed", d1)
		case d2 := <-ch2:
			fmt.Println("operation - 2 completed", d2)
		}
	}
}
