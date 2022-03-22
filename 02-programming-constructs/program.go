package main

import "fmt"

func main() {
	//if else
	/*
		var no = 21
		if no%2 == 0 {
			fmt.Println(no, "is even")
		} else {
			fmt.Println(no, "is odd")
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Println(no, "is even")
	} else {
		fmt.Println(no, "is odd")
	}

	/* for */
	fmt.Println("for v1.0")
	for idx := 0; idx < 10; idx++ {
		fmt.Println(idx * 10)
	}

	fmt.Println("for v2.0 (while)")
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println("numSum =", numSum)

	fmt.Println("for v2.0 (infinite loop)")
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println("x =", x)
}
