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
