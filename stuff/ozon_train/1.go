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

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--

		var s string
		fmt.Fscan(in, &s)
		if len(s) == 1 {
			fmt.Fprintln(out, 0)
			continue
		}
		n := len(s)
		answer := n - 1
		for i := 0; i < n-1; i++ {
			if s[i] < s[i+1] {
				answer = i
				break
			}
		}
		fmt.Fprintln(out, s[:answer]+s[answer+1:])
	}
}
