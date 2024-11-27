package main

import (
	"fmt"
	"sort"
)

func ExampleClosure() { // простой пример замыкания
	counter := 0
	increment := func(value int) {
		counter += value
	}

	fmt.Println("before", counter)
	increment(5)
	fmt.Println("after", counter)
}

func CreateClosureObjects() (inc, dec func(v int), get func() int) {
	var counter int = 0
	intFunc := func(inc int) {
		counter += inc
	}
	decFunc := func(dec int) {
		counter -= dec
	}
	getFunc := func() int {
		return counter
	}
	return intFunc, decFunc, getFunc
}

func ExampleClosureObject() {
	inc, dec, get := CreateClosureObjects()
	inc(5)
	dec(10)
	fmt.Println(get())
}

func ExampleClosureInSort() {
	var names = []string{
		"Аня",
		"Витя",
		"Коля",
		"Наташа",
		"Ирина",
	}
	sort.Slice(names, func(i, j int) bool { // сортировка при помощи замыкания по первой букве
		return []rune(names[i])[0] < []rune(names[j])[0] // т.к. строки хранятся в байтах, то нельзя просто написать names[i][0]
	})
	fmt.Println(names)
}

func main() {
	ExampleClosure()
	fmt.Println()
	ExampleClosureObject()
	ExampleClosureInSort()
}
