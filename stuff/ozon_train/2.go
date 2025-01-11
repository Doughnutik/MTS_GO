package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	need := map[rune]string{
		'D': "M",
		'C': "M",
		'R': "C",
		'M': "CRD",
	}

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--

		var s string
		fmt.Fscan(in, &s)
		cur_cond := 'D'
		flag := true
		for _, c := range s {
			flag = false
			for _, value := range need[cur_cond] {
				if value == c {
					cur_cond = value
					flag = true
					break
				}
			}
			if !flag {
				break
			}
		}
		if flag && cur_cond == 'D' {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
