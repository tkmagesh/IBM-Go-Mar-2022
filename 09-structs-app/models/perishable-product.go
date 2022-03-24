package models

import "fmt"

type PerishableProduct struct {
	product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		product{id, name, cost, units, category},
		expiry,
	}
}

func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("%s, Expiry=%s", format(pp.product), pp.Expiry)
}

func ApplyDiscount(pp *PerishableProduct, discount float32) {
	applyDiscount(&(pp.product), 10)
}
