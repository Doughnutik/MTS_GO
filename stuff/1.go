package main

import "fmt"

type Engine interface {
	Start()
	Power() int
}

type Car struct {
	Engine
}

type SlowEngine struct{}
type FastEngine struct{}

func (se SlowEngine) Start() {
	fmt.Println("Slow Start")
}

func (se SlowEngine) Power() int {
	return 100
}

func (fe FastEngine) Start() {
	fmt.Println("Fast Start")
}

func (fe FastEngine) Power() int {
	return 500
}

func dosmth(a ...any) { // пример функции, которая принимает много any и что-то делает в зависимости от типа
	for _, v := range a {
		switch t := v.(type) {
		case string:
			fmt.Println("string", t)
		case int:
			fmt.Println("int", t)
		}
	}
}

func main() {
	slow := SlowEngine{}
	fast := FastEngine{}

	slowCar := Car{Engine: slow}
	fastCar := Car{Engine: fast}

	slowCar.Start()
	fmt.Println(fastCar.Power())

	dosmth("amogus", 5)
}
