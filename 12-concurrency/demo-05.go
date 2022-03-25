package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}
	var goRoutineCount = 0
	if count, ok := strconv.Atoi(os.Args[1]); ok != nil {
		goRoutineCount = 1
	} else {
		goRoutineCount = count
	}
	fmt.Println("main started")
	for i := 1; i <= goRoutineCount; i++ {
		wg.Add(1)
		go f1(i, wg)
	}
	f2()
	wg.Wait() //block the execution of this function until the wg counter becomes 0
	fmt.Println("main completed")
}

func f1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("f1[%d] invocation - started\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("f1[%d] invocation - completed\n", id)
}

func f2() {
	fmt.Println("f2 invoked")
}
