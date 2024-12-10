package main

import "fmt"

type mystruct struct {
	a int
}

func (a *mystruct) change() {
	a.a = 5
}

func main() {
	a := map[string]int{"aboba": 1, "amogus": 2}
	b := a

	for key, _ := range b {
		b[key] = 100
	}

	fmt.Println(b)
	fmt.Println(a)
}
