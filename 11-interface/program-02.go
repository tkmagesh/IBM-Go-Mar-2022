package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

//implementation of fmt.Stringer interface
func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (productPtr *Product) ApplyDiscount(discount float32) {
	//dereference
	//(*productPtr).Cost = (*productPtr).Cost * ((100 - discount) / 100)
	productPtr.Cost = productPtr.Cost * ((100 - discount) / 100)
}
func main() {
	pen := Product{100, "Pen", 10, 50, "Stationary"}
	fmt.Println(pen)
}
