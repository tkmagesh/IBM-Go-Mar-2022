package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products Products) Print() {
	fmt.Println(products)
}

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%s\n", product)
	}
	return result
}

func (products Products) IndexOf(product Product) int {
	for idx, prod := range products {
		if prod == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products Products) Filter(predicate func(Product) bool) Products {
	filteredProducts := Products{}
	for _, prod := range products {
		if predicate(prod) {
			filteredProducts = append(filteredProducts, prod)
		}
	}
	return filteredProducts
}

func (products Products) All(predicate func(product Product) bool) bool {
	for _, prod := range products {
		if !predicate(prod) {
			return false
		}
	}
	return true
}

func (products Products) Any(predicate func(product Product) bool) bool {
	for _, prod := range products {
		if predicate(prod) {
			return true
		}
	}
	return false
}

func main() {

	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	//fmt.Println(FormatProducts(products))
	//fmt.Println(products.Format())
	products.Print()

	fmt.Println("Index Of:")
	marker := Product{103, "Marker", 50, 20, "Utencil"}
	//fmt.Println("Index of marker = ", IndexOf(products, marker))
	fmt.Println("Index of marker = ", products.IndexOf(marker))

	fmt.Printf("\nIncludes:\n")
	//fmt.Println("Does marker exist? : ", Includes(products, marker))
	fmt.Println("Does marker exist? : ", products.Includes(marker))

	fmt.Printf("\nFilter:\n")
	fmt.Println("Stationary products:")
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	//stationaryProducts := Filter(products, stationaryProductPredicate)
	stationaryProducts := products.Filter(stationaryProductPredicate)
	//fmt.Println(stationaryProducts.Format())
	stationaryProducts.Print()

	fmt.Println("Costly Products:")
	constlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}
	//costlyProducts := Filter(products, constlyProductPredicate)
	costlyProducts := products.Filter(constlyProductPredicate)
	//fmt.Println(costlyProducts.Format())
	costlyProducts.Print()

	fmt.Printf("\nAll:\n")
	//fmt.Println("Are all products stationary products ? :", All(products, stationaryProductPredicate))
	fmt.Println("Are all products stationary products ? :", products.All(stationaryProductPredicate))

	fmt.Printf("\nAny:\n")
	//fmt.Println("Are there any stationary products ? :", Any(products, stationaryProductPredicate))
	fmt.Println("Are there any stationary products ? :", products.Any(stationaryProductPredicate))
}
