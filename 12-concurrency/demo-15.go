package main

import (
	"fmt"
	"time"
)

func main() {
	//get the generated prime numbers and print them

	ch := genPrimes()
	for {
		if primeNo, ok := <-ch; ok {
			fmt.Println(primeNo)
			continue
		}
		break
	}
	fmt.Println("Done")
}

func genPrimes() chan int {
	count := 30
	ch := make(chan int)
	go func() {
		no := 2
		for count >= 0 {
			if isPrime(no) {
				time.Sleep(300 * time.Millisecond)
				ch <- no
				count -= 1
			}
			no++
		}
		close(ch)
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
