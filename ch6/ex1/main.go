package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func makePerson(firstName, lastName string, age int) Person {
	return Person{
		firstName,
		lastName,
		age,
	}
}

func makePersonPointer(firstName, lastName string, age int) *Person {
	personPtr := Person{
		firstName,
		lastName,
		age,
	}

	return &personPtr
}

func main() {
	me := makePerson("Precious", "Akpesiri", 26)
	myPrt := makePersonPointer("Precious", "Akpesiri", 26)

	fmt.Println(me)
	fmt.Println(myPrt)
}
