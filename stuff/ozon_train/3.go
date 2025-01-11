package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

		var n int
		var a []int
		fmt.Fscan(in, &n)package main

		import (
			"bufio"
			"fmt"
			"os"
			"sort"
			"strings"
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
		
				var n int
				var a []string
				fmt.Fscan(in, &n)
				a = make([]string, n)
				for i := 0; i < n; i++ {
					fmt.Fscan(in, &a[i])
				}
				sort.Slice(a, func(i int, j int) bool {
					if a[i][0] == '-' && a[j][0] == '-' {
						if len(a[i]) != len(a[j]) {
							return len(a[i]) > len(a[j])
						}
						return a[i] > a[j]
					} else if a[i][0] != '-' && a[j][0] != '-' {
						if len(a[i]) != len(a[j]) {
							return len(a[i]) < len(a[j])
						}
						return a[i] < a[j]
					} else {
						return a[i][0] == '-'
					}
				})
				// fmt.Println(a)
				s := strings.Join(a, " ")
				s += "\n"
				input, _ := in.ReadString('\n')
				input, _ = in.ReadString('\n')
				if input == s {
					fmt.Fprintln(out, "YES")
				} else {
					fmt.Fprintln(out, "NO")
				}
			}
		}
		
		a = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		fmt.Println(n)
		fmt.Println(a)
		sort.Ints(a)
		s := ""
		for _, num := range a {
			s += strconv.Itoa(num) + " "
		}
		s = s[:len(s)-1]
		input, _ := in.ReadString('\n')
		input = input[:len(input)-1]
		if input == s {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
