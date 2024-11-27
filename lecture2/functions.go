package main

type operation func(a, b int) int

var operations = map[string]operation{
	"+": add,
	"-": sub,
	"*": mul,
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func calc(op string, array []int) int {
	if _, ok := operations[op]; !ok || len(array) == 0 {
		return 0
	}
	var result int
	if op == "*" {
		result = 1
	}
	for _, value := range array {
		result = operations[op](result, value)
	}
	return result
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	println(calc("-", array))
	println(calc("+", array))
	println(calc("*", array))

	println()
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		println(pos(i), neg(-i))
	}
}
