/* Error Handling */

package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("invalid argument(s). cannot divide by zero")

func main() {
	quotient, remainder, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("do not attempt to divide by zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong")
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient = %d, Remainder = %d\n", quotient, remainder)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("invalid argument(s). cannot divide by zero")
		return 0, 0, err
	}
	return x / y, x % y, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
