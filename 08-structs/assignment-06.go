package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     5,
		Units:    100,
		Category: "Stationary",
	}

	/*
		fmt.Println(pen.Id) //accessing fields using 'value'
		penPtr := &pen
		//fmt.Println((*penPtr).Id) //accessing fields using 'pointer'
		fmt.Println(penPtr.Id)
	*/

	fmt.Println(Format(pen))
	ApplyDiscount(&pen, 10)
	fmt.Println("After applying discount")
	fmt.Println(Format(pen))
}

func Format(p Product) string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(productPtr *Product, discount float32) {
	//dereference
	//(*productPtr).Cost = (*productPtr).Cost * ((100 - discount) / 100)
	productPtr.Cost = productPtr.Cost * ((100 - discount) / 100)
}
