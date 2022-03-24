package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
	Org       Organization
}

type Organization struct {
	Name string
	City string
}

func NewEmployee(id int, firstName, lastName string, salary float32, orgName string, orgCity string) Employee {
	return Employee{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Salary:    salary,
		Org: Organization{
			Name: orgName,
			City: orgCity,
		},
	}
}

func main() {
	emp := Employee{
		Id:        100,
		FirstName: "Magesh",
		LastName:  "Kuppan",
		Salary:    10000,
		Org: Organization{
			Name: "IBM",
			City: "Bengaluru",
		},
	}
	fmt.Println(emp.Org.City)
}
