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
		lastName:  "keerthy",
	}

	fmt.Println("emp5", emp5)

	emp6 := Employee{"sam", "anderson", 55, 6000}
	fmt.Println("first", emp6.firstName)
	fmt.Println("last", emp6.lastName)

	var emp7 Employee
	emp7.firstName = "thomas"
	fmt.Println("emp7", emp7)

	emp8 := &Employee{"jimmy", "choo", 55, 6000}
	fmt.Println("first", (*emp8).firstName)
	fmt.Println("last", (*emp8).lastName)

	fmt.Println("first", emp8.firstName)
	fmt.Println("last", emp8.lastName)

	type Person struct {
		string
		int
	}

	p := Person{"naveen", 50}
	fmt.Println(p)

	p.string = "Jose"
	p.int = 100

	fmt.Println(p)

	type Address struct {
		city, state string
	}

	type Person2 struct {
		name    string
		age     int
		address Address
	}

	var p1 Person2

	p1.name = "umesh"
	p1.age = 35
	p1.address = Address{
		city:  "bangalore",
		state: "KA",
	}

	fmt.Println("name", p1.name)
	fmt.Println("age", p1.age)
	fmt.Println("city", p1.address.city)
	fmt.Println("state", p1.address.state)

	//promoted fields
	type Person3 struct {
		name string
		age  int
		Address
	}

	var p3 Person3
	p3.name = "Naveen"
	p3.age = 50
	p3.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p3.name)
	fmt.Println("Age:", p3.age)
	fmt.Println("City:", p3.city)   //city is promoted field
	fmt.Println("State:", p3.state) //state is promoted field

	type name struct {
		firstName string
		lastName  string
	}

	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}
}
