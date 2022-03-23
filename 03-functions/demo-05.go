/* higher order functions - functions can be passed as arguments to other functions */

package main

import "fmt"

func main() {

	/*
		logAddOperation(100, 200)
		logSubtractOperation(100, 200)
	*/
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)

	multiply := func(x, y int) {
		fmt.Println("multiply result =", x*y)
	}
	logOperation(multiply, 100, 200)

	logOperation(func(x, y int) {
		fmt.Println("divide result =", x/y)
	}, 100, 20)
}

/*
func logAddOperation(x, y int) {
	fmt.Println("Before invocation")
	add(x, y)
	fmt.Println("After invocation")
}

func logSubtractOperation(x, y int) {
	fmt.Println("Before invocation")
	subtract(x, y)
	fmt.Println("After invocation")
}
*/

func logOperation(operation func(int, int), x, y int) {
	fmt.Println("Before invocation")
	operation(x, y)
	fmt.Println("After invocation")
}

func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}
