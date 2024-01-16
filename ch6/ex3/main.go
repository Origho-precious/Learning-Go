package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func makePerson(firstName, lastName string, age int) []Person {
	var persons []Person

	for i := 0; i < 10000000; i++ {
		persons = append(persons, Person{
			firstName,
			lastName,
			age,
		})
	}

	return persons
}

func main() {

	start := time.Now()
	me := makePerson("Precious", "Akpesiri", 26)

	end := time.Now()

	fmt.Println("The call took", end.Sub(start))
	fmt.Println(len(me))
}
