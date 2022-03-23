package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(noPtr)

	fmt.Println(no == *(&no))
	//dereferencing (accessing the value using the pointer)
	/*
		var no2 = *noPtr
		fmt.Println(no2)
	*/
	fmt.Println("Before incrementing, no = ", no)
	increment(nil)
	fmt.Println("After incrementing, no = ", no)

	x, y := 10, 20
	fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)
}

func increment(x *int) {
	if x == nil {
		return
	}
	*x++
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}
