package main

import (
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

func mapIndex(k int, length int) int {
	var ki int
	if k < 0 {
		ki = length + k
	} else {
		ki = k
	}
	return ki
}

func shortestEdit(a string, b string) int {
	aLines := breakup(a)
	bLines := breakup(b)

	n, m := len(aLines), len(bLines)
	max := n + m
	//fmt.Printf("maximum number of moves for \n%s\nto \n%s\nis %v.\n", a, b, max)
	v := make([]int, 2*max+1) //long enough to hold values of x for any k
	v[1] = 0

	var fewestEdits int
	for d := 0; d <= max; d++ {
		// fmt.Printf("D = %v\n", d)
		for k := -d; k <= d; k += 2 { //changed k range to account for differences in array indexing
			//	fmt.Printf("  K = %v\n", k)

			var x int
			if k == -d || (k != d && v[mapIndex(k-1, len(v))] < v[mapIndex(k+1, len(v))]) { //move downward
				ki := mapIndex(k+1, len(v))
				x = v[ki]
			} else { //move rightward
				ki := mapIndex(k-1, len(v))
				x = v[ki] + 1
			}
			y := x - k
			//	fmt.Printf("(%v, %v)\n", x, y)
			for x < n && y < m && aLines[x].text == bLines[y].text { // diagonal move; represents deleting and inserting the same line
				x, y = x+1, y+1
			}
			ki := mapIndex(k, len(v))
			v[ki] = x
			//	fmt.Printf("d: %v | %v\n", d, v)
			if x >= n && y >= m {
				fewestEdits = d
				return fewestEdits
			}
		}
	}
	return fewestEdits
}

func main() {
	/*
		a := "ABCABBA"
		b := "CBABAC"
		fmt.Printf("string a: %s \nstring b: %s \n", a, b)
		aLines := breakup(a)
		bLines := breakup(b)
		fmt.Printf("lines of a: %v \nlines of b: %v \n", aLines, bLines)
		if strings.Compare(a, b) == 1 {
			fmt.Printf("strings a and b match\n")
		}
	*/
}
