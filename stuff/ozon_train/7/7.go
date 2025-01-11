package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var field [][]byte
var A, B [2]int

func fill_up(x int, y int, fl bool, ord bool) {
	if ord {
		for i := 0; i < y; i++ {
			if fl {
				field[x][i] = 'b'
			} else {
				field[x][i] = 'a'
			}
		}
		for i := 0; i < x; i++ {
			if fl {
				field[i][0] = 'b'
			} else {
				field[i][0] = 'a'
			}
		}
	} else {
		for i := 0; i < x; i++ {
			if fl {
				field[i][y] = 'b'
			} else {
				field[i][y] = 'a'
			}
		}
		for i := 0; i < y; i++ {
			if fl {
				field[0][i] = 'b'
			} else {
				field[0][i] = 'a'
			}
		}
	}
}

func fill_down(x int, y int, fl bool, ord bool) {
	if ord {
		for i := y + 1; i < m; i++ {
			if fl {
				field[x][i] = 'a'
			} else {
				field[x][i] = 'b'
			}
		}
		for i := x + 1; i < n; i++ {
			if fl {
				field[i][m-1] = 'a'
			} else {
				field[i][m-1] = 'b'
			}
		}
	} else {
		for i := x + 1; i < n; i++ {
			if fl {
				field[i][y] = 'a'
			} else {
				field[i][y] = 'b'
			}
		}
		for i := y + 1; i < m; i++ {
			if fl {
				field[n-1][i] = 'a'
			} else {
				field[n-1][i] = 'b'
			}
		}
	}
}

func search_up(fl bool) {
	if A[0] == 0 || field[A[0]-1][A[1]] == '#' {
		fill_up(A[0], A[1], fl, true)
	} else {
		fill_up(A[0], A[1], fl, false)
	}
}

func search_down(fl bool) {
	if B[0] == n-1 || field[B[0]+1][B[1]] == '#' {
		fill_down(B[0], B[1], fl, true)
	} else {
		fill_down(B[0], B[1], fl, false)
	}
}

func search() {
	fl := false
	if A[0] > B[0] || (A[0] == B[0] && A[1] > B[1]) {
		fl = true
		A[0], A[1], B[0], B[1] = B[0], B[1], A[0], A[1]
	}
	search_up(fl)
	search_down(fl)
}

func debug() {

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%c", field[i][j])
		}
		fmt.Println()
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--

		fmt.Fscan(in, &n, &m)
		field = make([][]byte, n)
		for i := 0; i < n; i++ {
			field[i] = make([]byte, m)
			var s string
			fmt.Fscan(in, &s)
			for j := 0; j < m; j++ {
				field[i][j] = s[j]
			}
		}
		// debug()
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if field[i][j] == 'A' {
					A = [2]int{i, j}
				} else if field[i][j] == 'B' {
					B = [2]int{i, j}
				}
			}
		}
		search()
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fprintf(out, "%c", field[i][j])
			}
			fmt.Fprintln(out)
		}
	}
}
