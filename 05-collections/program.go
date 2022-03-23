/* collections */

package main

import "fmt"

func main() {

	//Arrays

	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos [5]int = [...]int{3, 1, 4, 2, 5}
	nos := [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iteration using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	fmt.Println("Iteration using range")
	/*
		for idx, val := range nos {
			fmt.Printf("nos[%d] = %d\n", idx, val)
		}
	*/
	for _, val := range nos {
		fmt.Println(val)
	}
}
