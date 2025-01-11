package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var n, m int
var arr [][2]int
var cars [][4]int

func build() {
	sort.Slice(arr, func(i int, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	sort.Slice(cars, func(i int, j int) bool {
		if cars[i][0] != cars[j][0] {
			return cars[i][0] < cars[j][0]
		}
		return cars[i][3] < cars[j][3]
	})
}

func search() []int {
	answer := make([]int, n)
	load := make([]int, m)
	j := 0
	for i := 0; i < n; i++ {
		for j < m && (load[j] == cars[j][2] || arr[i][0] > cars[j][1]) {
			j++
		}
		if j == m || cars[j][0] > arr[i][0] {
			answer[arr[i][1]] = -1
		} else {
			load[j]++
			answer[arr[i][1]] = cars[j][3] + 1
		}
	}
	return answer
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

		fmt.Fscan(in, &n)
		arr = make([][2]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i][0])
			arr[i][1] = i
		}

		fmt.Fscan(in, &m)
		cars = make([][4]int, m)
		for i := 0; i < m; i++ {
			for j := 0; j < 3; j++ {
				fmt.Fscan(in, &cars[i][j])
			}
			cars[i][3] = i
		}

		build()
		answer := search()
		for _, i := range answer {
			fmt.Fprint(out, i, " ")
		}
		fmt.Fprintln(out)
	}
}
