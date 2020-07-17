package main

import (
	"fmt"
	"strings"
)

type line struct {
	number int
	text   string
}

func breakup(s string) []line {
	lines := make([]line, 0)
	if strings.Contains(s, "\n") {
		unnumbered := strings.Split(s, "\n")
		for i, v := range unnumbered {
			lin := line{i, v}
			lines = append(lines, lin)
		}
	} else {
		for i, c := range s {
			lin := line{i, string(c)}
			lines = append(lines, lin)
		}
	}
	return lines
}

func shortestEdit(a string, b string) int {
	aLines := breakup(a)
	bLines := breakup(b)

	n, m := len(a), len(b)
	max := n + m

	v := make([]int, 2*max+1)
	v[0] = 0

	var fewestEdits int
	for d := 0; d < max; d++ {
		for k := -d; k < d; k += 2 {
			var x int
			if k == -d || (k != d && v[k-1] < v[k+1]) {
				x = v[k+1]
			} else {
				x = v[k-1] + 1
			}
			y := x - k

			for x < n && y < m && aLines[x].text == bLines[y].text {
				x, y = x+1, y+1
			}
			v[k] = x

			if x >= n && y >= m {
				fewestEdits = d
			}
		}
	}
	return fewestEdits
}

func main() {
	a := "ABCABBA"
	b := "CBABAC"
	fmt.Printf("string a: %s \nstring b: %s \n", a, b)
	aLines := breakup(a)
	bLines := breakup(b)
	fmt.Printf("lines of a: %v \nlines of b: %v \n", aLines, bLines)
	if strings.Compare(a, b) == 1 {
		fmt.Printf("strings a and b match\n")
	}
}
