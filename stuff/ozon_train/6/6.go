package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF int = 1e9

var n, m int
var graph [][][2]int
var used []bool
var answer []int
var cnt int

func dfs(v int) bool {
	used[v] = true
	cur := false
	for _, u := range graph[v] {
		if !used[u[0]] {
			fl := dfs(u[0])
			if !fl {
				answer = append(answer, u[1])
				cur = !cur
			}
			cnt++
		}
	}
	return cur
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--

		fmt.Fscan(in, &n, &m)
		graph = make([][][2]int, n)
		for i := 0; i < m; i++ {
			var a, b int
			fmt.Fscan(in, &a, &b)
			a--
			b--
			graph[a] = append(graph[a], [2]int{b, i})
			graph[b] = append(graph[b], [2]int{a, i})
		}
		used = make([]bool, n)
		answer = []int{}
		cnt = 0
		for v := 0; v < n; v++ {
			if !used[v] {
				fl := dfs(v)
				if fl {
					cnt++
				}
			}
		}
		fmt.Fprintln(out, cnt)
		fmt.Fprintln(out, len(answer))
		for _, i := range answer {
			fmt.Fprint(out, i+1, " ")
		}
		fmt.Fprintln(out)
	}
}
