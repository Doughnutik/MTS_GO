package main

import (
	"fmt"

	rb "github.com/emirpasic/gods/trees/redblacktree"
)

var empt struct{}

var bar [][2]int
var pow [][2]int
var bucks [][]int
var n, m, L int

func myComp(a, b interface{}) int {
	x := a.([2]int)
	y := b.([2]int)
	for i := 0; i < 2; i++ {
		if x[i] < y[i] {
			return -1
		}
		if x[i] > y[i] {
			return 1
		}
	}
	return 0
}

func buildBucks() {
	ind := 0
	bucks = make([][]int, n)
	for i := 0; i < m; i++ {
		for ind < n && pow[i][0] > bar[ind][0] {
			ind++
		}
		if ind == n {
			break
		}
		bucks[ind] = append(bucks[ind], pow[i][1])
	}
}

func search() int {
	buildBucks()
	tree := rb.NewWith(myComp)
	answer := 0
	cnt := 0
	ma := 0
	curval := 1
	for i := 0; i < n; i++ {
		for curcnt, j := range bucks[i] {
			tree.Put([2]int{j, cnt + curcnt}, empt)
		}
		cnt += len(bucks[i])
		ma = max(ma, bar[i][1]-bar[i][0]+1)

		for !tree.Empty() && curval < ma+1 {
			key := tree.Right().Key
			tree.Remove(key)
			curval += key.([2]int)[0]
			answer++
		}
		if tree.Empty() && curval < ma+1 {
			return -1
		}
	}
	return answer
}

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		fmt.Scan(&n, &m, &L)
		bar = make([][2]int, n)
		pow = make([][2]int, m)
		for i := 0; i < n; i++ {
			fmt.Scan(&bar[i][0], &bar[i][1])
		}
		for i := 0; i < m; i++ {
			fmt.Scan(&pow[i][0], &pow[i][1])
		}
		fmt.Println(search())
		t--
	}
}
