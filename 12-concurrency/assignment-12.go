package main

import "fmt"

func main() {
	//get the generated prime numbers and print them
	count := 20
	ch := genPrimes(count)
	for i := 0; i < count; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func genPrimes(count int) chan int {
	ch := make(chan int)
	go func() {
		for count >= 0 {
			no := 2
			for {
				if isPrime(no) {
					ch <- no
					count -= 1
				}
				no++
			}
		}
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= no/2; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
