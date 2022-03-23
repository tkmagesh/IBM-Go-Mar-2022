/* collections */

package main

import "fmt"

func main() {

	//Arrays

	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos [5]int = [...]int{3, 1, 4, 2, 5}
	nos := [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iteration using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	fmt.Println("Iteration using range")
	/*
		for idx, val := range nos {
			fmt.Printf("nos[%d] = %d\n", idx, val)
		}
	*/
	for _, val := range nos {
		fmt.Println(val)
	}

	//slice
	fmt.Printf("\nSlice\n")
	//var products []string
	//var products []string = []string{"Pen", "Pencil", "Marker"}
	products := []string{"Pen", "Pencil", "Marker"}
	fmt.Println(products)
	fmt.Println("append")

	//adding 1 item
	products = append(products, "Scribble-pad")

	//adding more than 1 item
	products = append(products, "Mouse", "Charger")

	//adding items from another slice
	anotherList := []string{"ipad", "macbook"}
	products = append(products, anotherList...)

	fmt.Println(products)

	//slicing
	/*
		[lo:hi] => from lo to hi-1
		[lo:] => from lo to the end of the list
		[:hi] => from 0 to hi-1
		[:] => copy of the slice
	*/

	fmt.Println("products[2:5] => ", products[2:5])
	fmt.Println("products[:5] => ", products[:5])
	fmt.Println("products[5:] => ", products[5:])

	/*
		productsCopy := products[:]
	*/
	var productsCopy []string = make([]string, len(products))
	copy(productsCopy, products)
	products[0] = "mechanical-pen"
	fmt.Printf("products = %v, productsCopy = %v\n", products, productsCopy)

	fmt.Printf("\nMaps\n")
	//var productRanks map[string]int
	/*
		var productRanks map[string]int = make(map[string]int)
		//productRanks := map[string]int{}
		productRanks["Pen"] = 2
	*/
	//productRanks := map[string]int{"Pen": 2, "Pencil": 1, "Marker": 4, "Scribble-pad": 3}
	productRanks := map[string]int{
		"Pen":          2,
		"Pencil":       1,
		"Marker":       4,
		"Scribble-pad": 3,
	}
	fmt.Println(productRanks)

	fmt.Println("Iteration over a map")
	for key, value := range productRanks {
		fmt.Printf("Key = %s, Value = %d\n", key, value)
	}

	fmt.Printf("\nAdding a new item\n")
	productRanks["Mouse"] = 7
	fmt.Println(productRanks)

	fmt.Printf("\nAccessing a value by key\n")
	fmt.Printf("Rank of Mouse is : %d\n", productRanks["Mouse"])

	fmt.Printf("\nChecking for the existence of a key\n")
	keyToCheck := "Stylus"
	if val, exists := productRanks[keyToCheck]; !exists {
		fmt.Printf("Product %q does not exist\n", keyToCheck)
	} else {
		fmt.Printf("Rank of product %q is %d\n", keyToCheck, val)
	}

	fmt.Printf("\nRemoving a key\n")
	keyToRemove := "Stylus"
	delete(productRanks, keyToRemove)
	fmt.Printf("After deleting %q, productRanks = %v\n", keyToRemove, productRanks)

}
