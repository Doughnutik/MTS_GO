package main

import (
	"fmt"
	"time"
)

func fib(n int, mod int) int {
	results := make([]int, n+1)
	results[0] = 1
	results[1] = 1
	var recursive func(n int) // сначала нужно определить var
	recursive = func(n int) {
		if n < 2 {
			return
		}
		
		recursive(n - 1)
		results[n] = (results[n-1] + results[n-2]) % mod
	}
	recursive(n)
	// for _, i := range results {
	// 	fmt.Println(i)
	// }
	return results[n]
}

func main() {
	ts := time.Now()
	res := fib(1_000_000, 1e9+7) // можно так писать длинные числа
	duration := time.Since(ts)
	fmt.Println(res, duration)
}
