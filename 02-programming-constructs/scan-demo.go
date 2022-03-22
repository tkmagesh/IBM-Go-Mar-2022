package main

import "fmt"

func main() {
	var userName string
	fmt.Println("Enter your name :")
	//fmt.Scanln(&userName)
	fmt.Scanf("%s\n", &userName)
	//fmt.Println("Hi ", userName, ",Have a nice day")
	fmt.Printf("Hi %q, Have a nice day!\n", userName)

	var n1, n2 int
	fmt.Println("Enter 2 numbers")
	fmt.Scanln(&n1, &n2)
	//fmt.Scanf("%d %d\n", &n1, &n2)
	fmt.Println(n1, n2)

}
