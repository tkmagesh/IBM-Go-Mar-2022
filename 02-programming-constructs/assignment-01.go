package main

import "fmt"

func main() {
	for i := 3; i <= 100; i++ {
		var is_prime = true
		for n := 2; n <= (i / 2); n++ {
			if i%n == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			fmt.Println(i, " is prime")
		}
	}
}
