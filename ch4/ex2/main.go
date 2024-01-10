package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		randNum := rand.Intn(100)
		nums = append(nums, randNum)
	}

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i]%2 == 0 && nums[i]%3 == 0 {
	// 		fmt.Println("Six!")
	// 	} else if nums[i]%2 == 0 {
	// 		fmt.Println("Two!")
	// 	} else if nums[i]%3 == 0 {
	// 		fmt.Println("Three!")
	// 	} else {
	// 		fmt.Println("Never mind")
	// 	}
	// }

	// for i := 0; i < len(nums); i++ {
	// 	switch {
	// 	case nums[i]%2 == 0 && nums[i]%3 == 0:
	// 		fmt.Println("Six!")
	// 	case nums[i]%2 == 0:
	// 		fmt.Println("Two!")
	// 	case nums[i]%3 == 0:
	// 		fmt.Println("Three!")
	// 	default:
	// 		fmt.Println("Never mind")
	// 	}
	// }

	for _, v := range nums {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
