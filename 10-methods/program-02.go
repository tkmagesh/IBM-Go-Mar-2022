package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (productPtr *Product) ApplyDiscount(discount float32) {
	productPtr.Cost = productPtr.Cost * ((100 - discount) / 100)
}

//composition
type PerishableProduct struct {
	Product
	Expiry string
}

//method overriding
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry=%s", pp.Product.Format(), pp.Expiry)
}

func main() {
	pen := &Product{Id: 200, Name: "Pen", Cost: 10, Units: 40, Category: "Stationary"}

	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())

	//applying discount
	fmt.Println("After applying 20% discount")
	//ApplyDiscount(&pen, 20)
	//(&pen).ApplyDiscount(20)
	pen.ApplyDiscount(20)
	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())

	fmt.Printf("\nMethod inheritence\n")
	grapes := PerishableProduct{Product{200, "Grapes", 40, 60, "Fruits"}, "2 Days"}
	fmt.Println(grapes.Format())

	fmt.Println("After applying 20% discount")
	grapes.ApplyDiscount(20)
	fmt.Println(grapes.Format())
}
