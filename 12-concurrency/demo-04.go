package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var goRoutineCount = 0
	if count, ok := strconv.Atoi(os.Args[1]); ok != nil {
		goRoutineCount = 1
	} else {
		goRoutineCount = count
	}
	fmt.Println("main started")
	for i := 1; i <= goRoutineCount; i++ {
		wg.Add(1)
		go f1(i)
	}
	f2()
	wg.Wait() //block the execution of this function until the wg counter becomes 0
	fmt.Println("main completed")
}

func f1(id int) {
	fmt.Printf("f1[%d] invocation - started\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("f1[%d] invocation - completed\n", id)
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
