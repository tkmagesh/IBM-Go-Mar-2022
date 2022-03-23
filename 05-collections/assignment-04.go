package main

import "fmt"

func main() {
	result := generatePrimes(3, 100)
	fmt.Println(result)
}

func generatePrimes(start, end int) []int {
	result := make([]int, 0, 0)
	for i := 3; i <= 100; i++ {
		var is_prime = true
		for n := 2; n <= (i / 2); n++ {
			if i%n == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			result = append(result, i)
		}
	}
	return result
}
