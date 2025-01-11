package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func CheckFiles(files []string) int {
	cnt := 0
	for _, f := range files {
		a := strings.Split(f, ".")
		if a[len(a)-1] == "hack" {
			cnt++
		}
	}
	return cnt
}

func Parse(folder Folder) [2]int {
	cnt := [2]int{0, 0}
	cnt[0] += len(folder.Files)
	viruses := CheckFiles(folder.Files)
	for _, f := range folder.Folders {
		res := Parse(f)
		cnt[0] += res[0]
		if viruses > 0 {
			cnt[1] += res[0]
		} else {
			cnt[1] += res[1]
		}
	}
	if viruses > 0 {
		cnt[1] += len(folder.Files)
	}
	return cnt
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

		var n int
		fmt.Fscan(in, &n)
		sb := strings.Builder{}
		s, _ := in.ReadString('\n')
		for i := 0; i < n; i++ {
			s, _ = in.ReadString('\n')
			sb.WriteString(s)
		}
		var root Folder
		if err := json.Unmarshal([]byte(sb.String()), &root); err != nil {
			panic("Не смогли обработать data: " + err.Error())
		}
		res := Parse(root)
		fmt.Fprintln(out, res[1])
	}
}
