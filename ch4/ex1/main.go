package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// var nums []int = []int{}
	nums := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		randNum := rand.Intn(100)
		nums = append(nums, randNum)

		// fmt.Println(randNum)
	}

	fmt.Println(nums)
}
