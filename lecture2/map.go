package main

import "fmt"

func main() {
	var map1 map[int]int // danger, т.к. это nil map => записать в неё не получится
	println(len(map1))
	// map1[5] = 10    panic: assignment to entry in nil map
	map2 := make(map[int]int) // заранее можно указать cap
	map2[5] = 10
	for key, value := range map2 {
		println(key, value)
	}

	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	if val, ok := people["Tom"]; ok {
		fmt.Println(val)
	}
	if val, ok := people["Aboba"]; ok {
		fmt.Println(val) // новое значение в мапе не создастся
	}
	println("len:", len(people))
	// Если значение по заданному ключу имеется в отображении, то переменная ok будет равна true, а переменная val будет хранить полученное значение. Если переменная ok равна false, то значения по ключу в отображении нет.

	delete(people, "Tom")
	for key, value := range people {
		println(key, value)
	}
}
