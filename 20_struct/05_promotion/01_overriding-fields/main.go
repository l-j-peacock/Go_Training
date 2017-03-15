package main

import "fmt"

type Person struct {
	First	string
	Last 	string
	Age 	int
}

type DoubleZero struct {
	Person
	First 		string
	LicenceToKill	bool
}

func main() {
	p1 := DoubleZero {
		Person: Person {
			First: "James",
			Last: "Bond",
			Age: 20,
		},
		First: "Double Zero Seven",
		LicenceToKill: true,
	}

	p2 := DoubleZero {
		Person: Person {
			First: "Miss",
			Last: "Moneypenny",
			Age: 18,
		},
		First: "If looks could kill",
		LicenceToKill: false,
	}

	fmt.Println(p1.First, p1.Person.First)
	fmt.Println(p2.First, p2.Person.First)
}