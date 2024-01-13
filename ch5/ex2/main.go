package main

import (
	"fmt"
	"io"
	"os"
)

func fileLen(path string) (int, error) {
	file, err := os.Open(path)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	content, readErr := io.ReadAll(file)

	if readErr != nil {
		return 0, readErr
	}

	byteCount := len(content)

	return byteCount, nil

}

func main() {
	count, err := fileLen("../../ch4/ex2/main.go")

	fmt.Println("Count: ", count)
	fmt.Println("Err: ", err)

}
