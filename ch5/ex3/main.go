package main

import "fmt"

func prefixer(prefix string) func(prefix string) string {
	return func(text string) string {
		return prefix + " " + text
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
