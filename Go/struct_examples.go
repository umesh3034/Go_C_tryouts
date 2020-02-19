package main

import "fmt"

//named struct
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

	emp2 := Employee{"thomas", "paul", 29, 800}

	fmt.Println("emp1", emp1)
	fmt.Println("emp2", emp2)

	//anonymous struct
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "umesh",
		lastName:  "keerthy",
		age:       38,
		salary:    5000,
	}

	fmt.Println("emp3", emp3)

	//zero valued structure
	var emp4 Employee
	fmt.Println("emp4", emp4)

	emp5 := Employee{
		firstName: "umesh",
		lastName: ,
	}
}