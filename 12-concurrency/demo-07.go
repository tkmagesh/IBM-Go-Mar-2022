package main

import (
	"fmt"
	"sync"
)

var counter = 0 //race
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("counter = ", counter)
	fmt.Println("main completed")
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	{
		counter = counter + 1
	}
	mutex.Unlock()
}
