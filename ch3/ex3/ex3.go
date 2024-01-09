package main

import "fmt"

type Employee struct {
	id         int
	firstName  string
	secondName string
}

func main() {
	painter := Employee{0, "Pablo", "Picasso"}

	artist := Employee{
		id:         1,
		firstName:  "Yoyo",
		secondName: "Ma",
	}

	var scientist Employee = Employee{}

	scientist.id = 2
	scientist.firstName = "Isaac"
	scientist.secondName = "Newton"

	fmt.Println(painter)
	fmt.Println(artist)
	fmt.Println(scientist)
}
