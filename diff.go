package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
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
			lin := line{i + 1, v}
			lines = append(lines, lin)
		}
	} else {
		for i, c := range s {
			lin := line{i + 1, string(c)}
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

func shortestEdit(aLines []line, bLines []line) ([][]int, int) {
	n, m := len(aLines), len(bLines)
	max := n + m
	v := make([]int, 2*max+1) //long enough to hold values of x for any k
	v[1] = 0
	trace := make([][]int, 0)

	var fewestEdits int
	for d := 0; d <= max; d++ {
		vcopy := append([]int{}, v...)
		trace = append(trace, vcopy)
		for k := -d; k <= d; k += 2 { //changed k range to account for differences in array indexing
			var x int
			if k == -d || (k != d && v[mapIndex(k-1, len(v))] < v[mapIndex(k+1, len(v))]) { //move downward
				ki := mapIndex(k+1, len(v))
				x = v[ki]
			} else { //move rightward
				ki := mapIndex(k-1, len(v))
				x = v[ki] + 1
			}
			y := x - k
			for x < n && y < m && aLines[x].text == bLines[y].text { // diagonal move; represents deleting and inserting the same line
				x, y = x+1, y+1
			}
			ki := mapIndex(k, len(v))
			v[ki] = x
			if x >= n && y >= m {
				fewestEdits = d
				return trace, fewestEdits
			}
		}
	}
	return trace, fewestEdits
}

type move struct {
	px, py, x, y int
}

func backtrack(aLines []line, bLines []line) []move {
	x, y := len(aLines), len(bLines)
	moves := []move{}
	for trace, d := shortestEdit(aLines, bLines); d >= 0; d-- {
		v := trace[d]
		k := x - y
		var pk int
		if k == -d || (k != d && v[mapIndex(k-1, len(v))] < v[mapIndex(k+1, len(v))]) {
			pk = k + 1
		} else {
			pk = k - 1
		}
		px := v[pk]
		py := px - pk

		for x > px && y > py {
			fmt.Printf("(%v, %v) -> (%v, %v) \n", x-1, y-1, x, y)
			moves = append(moves, move{x - 1, y - 1, x, y})
			x, y = x-1, y-1
		}

		if d > 0 {
			fmt.Printf("(%v, %v) -> (%v, %v) \n", px, py, x, y)
			moves = append(moves, move{px, py, x, y})
		}
		x, y = px, py
	}
	return moves
}

type edit struct {
	kind    int
	oldLine line
	newLine line
}

func diff(aLines []line, bLines []line) []edit {
	diff := make([]edit, 0)
	var aLine, bLine line
	for _, move := range backtrack(aLines, bLines) {
		if move.px != len(aLines) {
			aLine = aLines[move.px]
		}
		if move.py != len(bLines) {
			bLine = bLines[move.py]
		}

		if move.x == move.px {
			new := edit{kind: 1, newLine: bLine}
			diff = append(diff, new)
		} else if move.y == move.py {
			new := edit{kind: -1, oldLine: aLine}
			diff = append(diff, new)
		} else {
			new := edit{kind: 0, oldLine: aLine, newLine: bLine}
			diff = append(diff, new)
		}
	}

	return diff //in reverse order
}

func print(diff []edit) {
	for i := len(diff) - 1; i >= 0; i-- {
		e := diff[i]
		if e.kind == 1 {
			color.Green("+	 	%v	%s\n", e.newLine.number, e.newLine.text)
			//fmt.Printf("+	 	%v	%s\n", e.newLine.number, e.newLine.text)
		} else if e.kind == -1 {
			color.Red("-	%v	 	%s\n", e.oldLine.number, e.oldLine.text)
			//fmt.Printf("-	%v	 	%s\n", e.oldLine.number, e.oldLine.text)
		} else {
			fmt.Printf(" 	%v	%v	%s\n", e.oldLine.number, e.newLine.number, e.oldLine.text)
		}
	}
}

func main() {

	a := "ABCABBA"
	b := "CBABAC"
	aLines, bLines := breakup(a), breakup(b)
	fmt.Printf("string a: %s \nstring b: %s \n", a, b)
	/*
		trace, _ := shortestEdit(a, b)
		for d, v := range trace {
			fmt.Printf("%v | %v\n", d, v)
		}
	*/
	//fmt.Printf("%v\n", backtrack(aLines, bLines))
	diff := diff(aLines, bLines)
	for i, e := range diff {
		fmt.Printf("%v: %v\n", i, e)
	}
	print(diff)
}
