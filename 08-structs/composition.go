package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Expiry string
	Name   string
}

func main() {
	/*
		grapes := PerishableProduct{
			Product{200, "Grapes", 20, 40, "Fruits"},
			"2 Days",
		} */

	grapes := PerishableProduct{
		Product: Product{Id: 200, Name: "Grapes", Cost: 20, Units: 40, Category: "Fruits"},
		Expiry:  "2 Days",
	}

	fmt.Println(grapes.Product.Name)
	fmt.Println(grapes.Name)

	//fmt.Println(Format(grapes.Product))
	fmt.Println(FormatPerishableProduct(grapes))
	fmt.Println("After applying discount")
	ApplyDiscount(&grapes.Product, 10)
	//fmt.Println(Format(grapes.Product))
	fmt.Println(FormatPerishableProduct(grapes))

}

func Format(p Product) string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(productPtr *Product, discount float32) {
	//dereference
	//(*productPtr).Cost = (*productPtr).Cost * ((100 - discount) / 100)
	productPtr.Cost = productPtr.Cost * ((100 - discount) / 100)
}

func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("%s, Expiry=%s", Format(pp.Product), pp.Expiry)
}
