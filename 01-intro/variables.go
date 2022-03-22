package main

import "fmt"

var x = 100

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	//type inference
	/*
		var msg = "Hello World!"
	*/

	msg := "Hello World!" // := syntax is application only in a function (not at package level)
	fmt.Println(msg)

	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Result of adding ="
		result = x + y
		fmt.Println(x, y, str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Result of addtion ="
		var result int = x + y
		fmt.Println(x, y, str, result)
	*/

	/* var x, y, result int
	var str string
	x = 100
	y = 200
	str = "Result of adding ="
	result = x + y
	fmt.Println(x, y, str, result) */

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Result of adding ="
		result = x + y
		fmt.Println(x, y, str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Result of addition ="
			result int    = x + y
		)
	*/

	/*
		x, y := 100, 200
		str := "Result of addition ="
		result := x + y
	*/

	x, y, str := 100, 200, "Result of addition ="
	result := x + y

	fmt.Println(x, y, str, result)

	//constants
	const pi = 3.14

	//iota
	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/
	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 3
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 3
			green
			_
			blue
		)
	*/

	const (
		red = iota * 3
		green
		_
		blue
	)
	fmt.Println("red =", red, "green =", green, "blue =", blue)
}
