/* function basics */

package main

import "fmt"

func main() {
	fn()
	greet("Magesh")
	fmt.Println(getGreetMsg("Suresh"))
	fmt.Printf("Adding 100 & 200, result = %d\n", add(100, 200))

	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d, remainder = %d\n", quotient, remainder)
	*/

	/* _ can be used in place of a varible that we dont want to use */
	quotient, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
}

func fn() {
	fmt.Println("fn invoked")
}

func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day\n", userName)
}

func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day\n", userName)
}

/*
func add(x int, y int) int {
	return x + y
}
*/

/*
func add(x, y int) int {
	return x + y
}
*/

func add(x, y int) (result int) {
	result = x + y
	return
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
