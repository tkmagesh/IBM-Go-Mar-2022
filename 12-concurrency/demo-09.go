package main

import (
	"fmt"
	"sync"
)

//Communicate By Sharing Memory -- DO NOT DO THIS (instead Share Memory for communication)
var result = 0

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println(result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	result = x + y
}
