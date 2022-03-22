/* variadic functions */
package main

import "fmt"

func main() {
	fmt.Println(sum("my string", 10))
	fmt.Println(sum("my string", 10, 20))
	fmt.Println(sum("my string", 10, 20, 30))
	fmt.Println(sum("my string", 10, 20, 30, 40, 50, 60))
}

func sum(s string, nos ...int) int {
	result := 0
	for idx := 0; idx < len(nos); idx++ {
		result += nos[idx]
	}
	return result
}
