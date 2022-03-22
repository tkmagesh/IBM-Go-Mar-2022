package main

import "fmt"

func main() {
	//if else
	/*
		var no = 21
		if no%2 == 0 {
			fmt.Println(no, "is even")
		} else {
			fmt.Println(no, "is odd")
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Println(no, "is even")
	} else {
		fmt.Println(no, "is odd")
	}

	/* for */
	fmt.Println("for v1.0")
	for idx := 0; idx < 10; idx++ {
		fmt.Println(idx * 10)
	}

	fmt.Println("for v2.0 (while)")
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println("numSum =", numSum)

	fmt.Println("for v3.0 (infinite loop)")
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println("x =", x)

	//switch case
	/*
		score 0-3 => "Terrible"
		score 4-7 => "Not Bad"
		score 8-9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	score := 7
	/* switch score {
	case 0:
		fmt.Println("Terrible")
	case 1:
		fmt.Println("Terrible")
	case 2:
		fmt.Println("Terrible")
	case 3:
		fmt.Println("Terrible")
	case 4:
		fmt.Println("Not Bad")
	case 5:
		fmt.Println("Not Bad")
	case 6:
		fmt.Println("Not Bad")
	case 7:
		fmt.Println("Not Bad")
	case 8:
		fmt.Println("Very Good")
	case 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	} */

	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	no := 15
	switch {
	case no%3 == 0:
		fmt.Println(no, "is divisible by 3")
	case no%4 == 0:
		fmt.Println(no, "is divisible by 4")
	}

	//falthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
	}

	subscription := "premium"
	switch subscription {
	case "premium":
		fmt.Println("All <premium> features are included")
		fallthrough
	case "advanced":
		fmt.Println("All <advanced> features are included")
		fallthrough
	case "basic":
		fmt.Println("All <basic> features are included")
		fallthrough
	default:
		fmt.Println("All <free> features are included")
	}
}
