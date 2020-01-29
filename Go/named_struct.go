package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	emp1 := Employee{
		firstName: "sam",
		age:       25,
		salary:    500,
		lastName:  "anderson",
	}

	emp2 := Employee{"thmomas", "paul", 29, 800}

	fmt.Println("Emp1:", emp1)
	fmt.Println("Emp2:", emp2)
}
