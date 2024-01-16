package main

import "fmt"

func updateSlice(strSlice []string, str string) {
	strSlice[len(strSlice)-1] = str
	fmt.Println("in updateSlice:", strSlice)
}

func growSlice(strSlice []string, str string) {
	strSlice = append(strSlice, str)
	fmt.Println("in growSlice:", strSlice)
}

func main() {
	var slice1 = []string{"me", "you", "mine", "yours"}
	var slice2 = []string{"me", "you", "mine", "yours"}

	updateSlice(slice1, "them")

	fmt.Println("in main after updateSlice:", slice1)

	growSlice(slice2, "them")

	fmt.Println("in main after growSlice:", slice2)
}
