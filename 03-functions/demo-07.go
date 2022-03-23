/* higher order functions - functions can be passed as arguments to other functions */

package main

import "fmt"

func main() {

	add(100, 200)
	subtract(100, 200)

	logOperation(add, 100, 200)
	logAdd(100, 200)

	logOperation(subtract, 100, 200)
	logSubtract(100, 200)
}

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
