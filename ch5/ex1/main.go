package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(i int, j int) (int, error) {
	return i + j, nil
}

func sub(i int, j int) (int, error) {
	return i - j, nil
}

func mul(i int, j int) (int, error) {
	return i * j, nil
}

func div(i int, j int) (int, error) {
	if i == 0 || j == 0 {
		err := errors.New("division by zero")

		return 0, err
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "0"},
		{"2", "%", "0"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])

		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]

		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])

		if err != nil {
			fmt.Println(err)
			continue
		}

		result, opErr := opFunc(p1, p2)

		if opErr != nil {
			fmt.Println(opErr)
			continue
		}

		fmt.Println(result)
	}
}
