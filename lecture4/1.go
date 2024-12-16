package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("input.txt")
	err2 := fmt.Errorf("tried to open file: %w", err)
	fmt.Println(err2)
}
