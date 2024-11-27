package main

import "fmt"

type Vehicle interface { // аналог темплейта в плюсах
	getPower() int
	getPrice() int
	move()
}

type Bus struct {
	power int
	price int
}

type Car struct {
	power int
	price int
}

func (car Car) getPower() int { // методы реализуются вне структуры
	return car.power
}

func (bus Bus) getPower() int {
	return bus.power
}

func (car Car) getPrice() int {
	return car.price
}

func (bus Bus) getPrice() int {
	return bus.price
}

func (car Car) move() {
	fmt.Println("Car is moving")
}

func (bus Bus) move() {
	fmt.Println("Bus is moving")
}

func main() {
	var new_car Vehicle = Car{100, 1000}
	var new_bus Vehicle = Bus{500, 10000}
	fmt.Println(new_car.getPower(), new_car.getPrice())
	new_car.move()
	fmt.Println(new_bus.getPower(), new_bus.getPrice())
	new_bus.move()
}
