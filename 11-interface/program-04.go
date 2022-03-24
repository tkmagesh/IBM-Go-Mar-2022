/* interface{} (empty interface) */
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type ShapeWithArea interface {
	Area() float32
}

func main() {
	var x interface{}
	//x = 100
	//x = "abc"
	//x = true
	//x = 10.000
	x = Circle{Radius: 10}

	if val, ok := x.(int); ok {
		fmt.Println(val + 200)
	} else {
		fmt.Println("Not an integer")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("x is integer, x + 100 = ", val+100)
	case string:
		fmt.Println("x is string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is a boolean, value of x =", val)
	case ShapeWithArea:
		fmt.Println("x is something with Area, Area = ", val.Area())
	default:
		fmt.Println("x is of unknown type")
	}

}
