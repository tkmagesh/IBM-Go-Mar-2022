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
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	x, y := 10, 20
	fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)
	swap( /*  */ )
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)
}

func increment(x *int) {
	*x++
}

func swap( /*  */ ) {

}
