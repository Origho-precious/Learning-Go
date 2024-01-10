package main

import "fmt"

func main() {
	var total int
	var array [10]int

	for i := range array {
		total := total + i // Change := to = to fix the bug
		fmt.Println(total)
	}
	fmt.Println(total)
}
