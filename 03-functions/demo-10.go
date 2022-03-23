/* Error Handling */

package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("invalid argument(s). cannot divide by zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred. exiting the application")
			fmt.Println(err)
		}
	}()
	quotient, remainder, err := divideClient(100, 0)
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

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("divideClient resulted in an error")
			err = DivideByZeroError
			return
		}
		fmt.Println("divideClient executed without any errors")
	}()
	quotient, remainder = divide(x, y)
	return
}

//library function
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
