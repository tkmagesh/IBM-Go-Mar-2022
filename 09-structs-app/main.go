package main

import (
	"fmt"
	"structs-app/models"
)

func main() {
	grapes := models.NewPerishableProduct(200, "Grapes", 20, 40, "Fruits", "2 Days")
	fmt.Println(grapes.Name)

	fmt.Println(models.FormatPerishableProduct(grapes))
	fmt.Println("After applying discount")
	models.ApplyDiscount(&grapes, 10)
	fmt.Println(models.FormatPerishableProduct(grapes))
}
