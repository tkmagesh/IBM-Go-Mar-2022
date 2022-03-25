package main

import (
	"fmt"
)

func main() {
	fmt.Println("main started")
	go f1() // DONOT wait for the completion of the function execution. instead schedule this function for execution
	f2()
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
