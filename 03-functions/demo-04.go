/* higher order functions - functions can be assigned as values to variables */
package main

import "fmt"

func main() {
	var greet func(string)
	/* fn := func() {
		fmt.Println("fn invoked")
	} */
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	/*
		greet := func(userName string) {
			fmt.Printf("Hi %s, Have a nice day!\n", userName)
		}
	*/

	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	/* isEven := func(no int) bool {
		return no%2 == 0
	} */

	var isEven func(int) bool
	isEven = func(no int) bool {
		return no%2 == 0
	}
	fmt.Printf("Is 21 even ? : %t\n", isEven(21))

	var operation func(int, int) int
	operation = func(x, y int) int {
		return x + y
	}
	fmt.Println(operation(100, 200))

	operation = func(x, y int) int {
		return x - y
	}
	fmt.Println(operation(100, 200))
}
