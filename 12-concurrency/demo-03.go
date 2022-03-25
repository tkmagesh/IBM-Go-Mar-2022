package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1() // DONOT wait for the completion of the function execution. instead schedule this function for execution
	f2()

	var input string
	fmt.Scanln(&input)

	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 invocation - started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 invocation - completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
