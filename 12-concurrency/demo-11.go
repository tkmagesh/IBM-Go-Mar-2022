package main

import "fmt"

func main() {
	//The following will result in a deadlock
	/*
		ch := make(chan int)
		ch <- 100
		val := <-ch
		fmt.Println(val) */

	//Solution - 1
	/*
		ch := make(chan int)
		go func(){
			ch <- 100
		}()
		val := <-ch
		fmt.Println(val)
	*/
	//Solution - 2 (using buffered channel)
	ch := make(chan int, 1)
	ch <- 100
	val := <-ch
	fmt.Println(val)
}
