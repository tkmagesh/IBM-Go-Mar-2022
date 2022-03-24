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

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

type ShapeWithArea interface {
	Area() float32
}

func PrintArea(sa ShapeWithArea) {
	fmt.Printf("Area = %f\n", sa.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(sp ShapeWithPerimeter) {
	fmt.Printf("Perimeter = %f\n", sp.Perimeter())
}

/*
	x => any type that implements BOTH shapeWithArea & shapeWithPerimeter interfaces
*/

//interface composition
type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintShape(x Shape) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{Radius: 12}
	//fmt.Printf("Area = %f\n", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 10, Width: 12}
	//fmt.Printf("Area = %f\n", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)

}
