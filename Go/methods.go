package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displaySalary() {
	fmt.Printf("salary %s %s %d\n", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee{"sam", 5000, "$"}
	emp1.displaySalary()
}
