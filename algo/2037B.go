package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var k int
		fmt.Scan(&k)
		a := make(map[int]int)
		for i := 0; i < k; i++ {
			var x int
			fmt.Scan(&x)
			a[x]++
		}
		for key := range a {
			if (k-2)%key != 0 {
				continue
			}
			x := (k - 2) / key
			if v, ok := a[x]; (x != key && ok) || (x == key && v >= 2) {
				fmt.Println(key, x)
				break
			}
		}
		t--
	}
}
