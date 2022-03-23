package main

import "fmt"

func main() {
	/*
		var nos []int
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))

		nos = append(nos, 3)
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))

		nos = append(nos, 5)
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))

		nos = append(nos, 2)
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))

		nos = append(nos, 1)
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))

		nos = append(nos, 4)
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))
	*/
	/*
		nos := []int{3, 1, 4, 2, 5}
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))

		nos = append(nos, 6)
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))
	*/
	/*
		nos := make([]int, 3, 5)
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))

		nos[0] = 3
		nos[1] = 1
		nos[2] = 4
		//nos[3] = 5  - throw an index out of range
		nos = append(nos, 5)
		fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))
	*/
	nos := make([]int, 0, 5)
	fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))

	nos = append(nos, 3)
	nos = append(nos, 1)
	nos = append(nos, 4)
	//nos[3] = 5  - throw an index out of range
	nos = append(nos, 5)
	fmt.Printf("nos = %v, len = %d, cap = %d\n", nos, len(nos), cap(nos))
}
