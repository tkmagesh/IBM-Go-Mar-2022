/* anonymous functions */

package main

import "fmt"

func main() {
	func() {
		fmt.Println("fn invoked")
	}()

	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

	quotient, remainder := func(x, y int) (int, int) {
		return x / y, x % y
	}(100, 7)
	fmt.Println(quotient, remainder)
}
