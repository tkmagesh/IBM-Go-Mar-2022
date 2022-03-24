package models

import "fmt"

type product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func format(p product) string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func applyDiscount(productPtr *product, discount float32) {
	//dereference
	//(*productPtr).Cost = (*productPtr).Cost * ((100 - discount) / 100)
	productPtr.Cost = productPtr.Cost * ((100 - discount) / 100)
}
