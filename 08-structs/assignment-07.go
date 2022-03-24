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

	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	fmt.Println(FormatProducts(products))

	fmt.Println("Index Of:")
	marker := Product{103, "Marker", 50, 20, "Utencil"}
	fmt.Println("Index of marker = ", IndexOf(products, marker))

	fmt.Printf("\nIncludes:\n")
	fmt.Println("Does marker exist? : ", Includes(products, marker))

	fmt.Printf("\nFilter:\n")
	fmt.Println("Stationary products:")
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := Filter(products, stationaryProductPredicate)
	fmt.Println(FormatProducts(stationaryProducts))

	fmt.Println("Costly Products:")
	constlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}
	costlyProducts := Filter(products, constlyProductPredicate)
	fmt.Println(FormatProducts(costlyProducts))

	fmt.Printf("\nAll:\n")
	fmt.Println("Are all products stationary products ? :", All(products, stationaryProductPredicate))

	fmt.Printf("\nAny:\n")
	fmt.Println("Are there any stationary products ? :", Any(products, stationaryProductPredicate))
}

func Format(p Product) string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func FormatProducts(products []Product) string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%s\n", Format(product))
	}
	return result
}

func IndexOf(products []Product, product Product) int {
	for idx, prod := range products {
		if prod == product {
			return idx
		}
	}
	return -1
}

func Includes(products []Product, product Product) bool {
	return IndexOf(products, product) != -1
}

func Filter(products []Product, predicate func(Product) bool) []Product {
	filteredProducts := []Product{}
	for _, prod := range products {
		if predicate(prod) {
			filteredProducts = append(filteredProducts, prod)
		}
	}
	return filteredProducts
}

func All(products []Product, predicate func(product Product) bool) bool {
	for _, prod := range products {
		if !predicate(prod) {
			return false
		}
	}
	return true
}

func Any(products []Product, predicate func(product Product) bool) bool {
	for _, prod := range products {
		if predicate(prod) {
			return true
		}
	}
	return false
}
