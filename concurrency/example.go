package main

import (
	"fmt"
	"time"
)

func main() {
	go test() // go - ключевое слово для функции, которая будет выполняться конкурентно/параллельно
	go func() {
		fmt.Println("closure function is running")
	}()

	// Здесь мы добавили 2 горутины, но функция main завершается быстрее, чем планировщик начнёт их исполнять.
	// Поэтому вывод пустой. Чтобы это починить, добавил sleep

	time.Sleep(1 * time.Second)

	// Заметим, что порядок выполнения функций неопределён.
	// Также важно, что планировщик сам решит, выполнять горутины конкурентно или параллельно. Поэтому 100% многопоточности здесь нет!
	// Причины, почему может быть конкурентное выполнение - отсутствие свободных ядер, обращение к одной памяти.
}

func test() {
	fmt.Println("test function is running")
}
