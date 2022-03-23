package main

import "fmt"

type operation func(int, int) int

var operations map[int]operation = map[int]operation{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	for {
		userChoice := getUserChoice()
		if userChoice == 5 {
			break
		}
		processUserChoice(userChoice)
	}
}

func processUserChoice(userChoice int) {
	if operation, exists := operations[userChoice]; exists {
		n1, n2 := getOperands()
		result := operation(n1, n2)
		fmt.Printf("Result = %d\n", result)
		return
	}
	fmt.Println("invalid choice")
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("Enter your choice")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Scanln(&userChoice)
	return userChoice
}

func getOperands() (int, int) {
	var n1, n2 int
	fmt.Println("Enter the nos :")
	fmt.Scanln(&n1, &n2)
	return n1, n2
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
