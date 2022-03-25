package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32 = 0 //race

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
	atomic.AddInt32(&counter, 1)
}
