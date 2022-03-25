package main

import (
	"fmt"
	"time"
)

func main() {
	//get the generated prime numbers and print them

	ch := genPrimes()
	for primeNo := range ch {
		fmt.Println(primeNo)
	}
	fmt.Println("Done")
}

//generate as many prime numbers as possible within 30 seconds time period
func genPrimes() chan int {
	ch := make(chan int)
	go func() {
		no := 2
		done := false
		go func() {
			time.Sleep(10 * time.Second)
			done = true
			fmt.Println(done)
		}()
		for !done {
			if isPrime(no) {
				time.Sleep(300 * time.Millisecond)
				ch <- no
			}
			no++
		}
		close(ch)
	}()
	return ch
}

func timeOut(d time.Duration) chan bool {
	timeOutCh := make(chan bool)
	go func() {
		time.Sleep(d)
		timeOutCh <- true
	}()
	return timeOutCh
}

func isPrime(no int) bool {
	for i := 2; i <= no/2; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
