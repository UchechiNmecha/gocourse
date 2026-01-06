package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	emp := Employee{
		FirstName: "Uceeey",
		LastName:  "Erudite",
		Age:       30,
	}

	const MAXENTRIES = 100

	fmt.Println("Employee:", emp.FirstName, emp.LastName, "Age:", emp.Age)
}