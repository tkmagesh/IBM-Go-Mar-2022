/* higher order functions - functions can be returned as return values from other functions */
package main

import "fmt"

func main() {
	greeter := getGreeter()
	greeter("Magesh")
	adderFor100 := getAdder(100)
	fmt.Println(adderFor100(200))
}

func getGreeter() func(string) {
	greeter := func(userName string) {
		fmt.Printf("Hi %s\n", userName)
	}
	return greeter
}

func getAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
