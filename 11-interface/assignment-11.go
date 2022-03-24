package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func Format(p Product) string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products Products) Print() {
	fmt.Println(products.Format())
}

func (products Products) Format() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%s\n", Format(product))
	}
	return result
}

//implement "sorting" for the products
//The user of the Products type should be able to sort the products by id / name / cost / units / category
//NOTE: Use sort.Sort() function

func main() {

}
