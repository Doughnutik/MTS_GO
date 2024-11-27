package main

import "fmt"

func main() {
	// Строки неизменяемы!!!
	a := "aboba"
	// Так сделать нельзя
	// a[0] = "a"
	// a[0] = 'a'
	b := "amogus"
	c := a + b // так можно
	println(a, b, c)

	d := "привет"
	println(len(d)) // длина вычисляется в байтах, а не в символах

	d = a[:1]
	println(d, d[0])      // как видим, выводится код символа, а не сам символ, потому что храним байты
	println(string(d[0])) // так выведем символ

	s := "hello"
	for i, rn := range s {
		fmt.Printf("%2v: 0x%x %v", i, rn, string(rn))
	}
}
