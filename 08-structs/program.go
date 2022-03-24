package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{100, "Magesh", "Kuppan", 10000}
	//var emp Employee = Employee{Id: 100, FirstName: "Magesh", Salary: 10000}
	emp := Employee{
		Id:        100,
		FirstName: "Magesh",
		Salary:    10000,
	}
	fmt.Printf("%#v\n", emp)

	e1 := Employee{
		Id:        200,
		FirstName: "Suresh",
		Salary:    2000,
	}

	e2 := Employee{
		Id:        200,
		FirstName: "Suresh",
		Salary:    2000,
	}

	fmt.Println(e1 == e2) //true
	fmt.Printf("%p %p\n", &e1, &e2)

	fmt.Println(true && true)
}
