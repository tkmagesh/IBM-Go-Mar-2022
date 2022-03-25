package main

import (
	"fmt"
	"sync"
)

//Share Memory for communication

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	/*
		<-  channel operator
		ch <- <data>   => sending operation (write)
		<- ch    => receiving operation (read)
	*/
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch
	fmt.Println(result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := x + y
	ch <- result
}
