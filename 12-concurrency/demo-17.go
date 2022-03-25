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

	//timeOutCh := timeOut(10 * time.Second)
	timeOutCh := time.After(10 * time.Second)
	primeNoCh := make(chan int)
	go func() {
		no := 2
	LOOP:
		for {
			if isPrime(no) {
				select {
				case <-timeOutCh:
					break LOOP
				case primeNoCh <- no:
					time.Sleep(300 * time.Millisecond)
				}
			}
			no++
		}
		close(primeNoCh)
	}()
	return primeNoCh
}

func timeOut(d time.Duration) chan time.Time {
	timeOutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeOutCh <- time.Now()
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
