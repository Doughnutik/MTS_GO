package main

import (
	"fmt"
	"sort"
)

func twoSum(a []int, target int) []int {
	n := len(a)
	b := make([][2]int, n)
	for i := 0; i < n; i++ {
		b[i][0] = a[i]
		b[i][1] = i + 1
	}
	sort.Slice(b, func(i int, j int) bool {
		return b[i][0] < b[j][0]
	})
	j := n - 1
	for i := 0; i < j; i++ {
		for j > i && b[i][0]+b[j][0] > target {
			j--
		}
		if i == j {
			panic("smth wrong")
		}
		if b[i][0]+b[j][0] == target {
			return []int{b[i][1], b[j][1]}
		}
	}
	return []int{}
}

func main() {
	qu := []int{2, 1}
	v, u := qu[0], qu[1]
	if v > u {
		v, u = u, v
	}
	fmt.Println(v, u)
}
