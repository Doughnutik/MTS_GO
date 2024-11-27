package main

import "fmt"

func f(a []int) {
	a[0] = 10
	fmt.Println(a)
	a = append(a, 100)
	fmt.Println(a)
	a[2] = 30
	fmt.Println(a)
}
