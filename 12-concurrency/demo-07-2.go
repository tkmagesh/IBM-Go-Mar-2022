package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.value++
	}
	c.Unlock()
}

func main() {

	wg := &sync.WaitGroup{}
	counter := &Counter{}
	fmt.Println("main started")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg, counter)
	}
	wg.Wait()
	fmt.Println("counter = ", counter.value)
	fmt.Println("main completed")
}

func fn(wg *sync.WaitGroup, counter *Counter) {
	defer wg.Done()
	counter.Increment()
}
