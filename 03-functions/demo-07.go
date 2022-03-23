/* higher order functions - functions can be passed as arguments to other functions */

package main

import "fmt"

func main() {

	add(100, 200)
	subtract(100, 200)

	//logOperation(add, 100, 200)
	logAdd := getLogOperation(add)
	logAdd(100, 200)

	//logOperation(subtract, 100, 200)
	logSubtract := getLogOperation(subtract)
	logSubtract(100, 200)
}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("Before invocation")
		operation(x, y)
		fmt.Println("After invocation")
	}
}

func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}
