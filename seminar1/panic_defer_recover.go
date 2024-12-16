package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func divideByZero() {
	// fmt.Println(divide(1, 0))
	// Важно, что defer должен быть перед местом потенциального возникновения паники. Если будет после, то не перехватит!
	defer func() { // отложенная функция
		if err := recover(); err != nil { // если случилась паника
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(divide(1, 0)) // сначала вызовется это выражение, затем defer func
}
func divide(a, b int) int {
	return a / b
}

func ExampleDeferFile() {
	fd, err := os.OpenFile("input.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Open error", err) // можно попробовать указать несуществующий файл. GO по умолчанию его не создаст
		return
	}
	defer fd.Close() // Это нужно, чтобы даже при возникновении ошибок чтений, дескриптор всё равно закрылся и не было никаких утечек/блокировок.
	// Если же файл так и не открылся (произашло то, что перед defer), то он и не будет закрываться.

	reader := bufio.NewReader(fd)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read error", err)
		return
	}
	fmt.Println("Read line", line)
}

func main() {
	divideByZero()
	fmt.Println("we survived dividing by zero!")
	ExampleDeferFile()
}
