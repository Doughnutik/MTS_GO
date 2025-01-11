package main

import (
	"fmt"
)

func main() {
	s := "Hello.for.after.back.hack"
	res := s.Split(".")
	fmt.Println(res)
}
