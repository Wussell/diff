package main

import "testing"

func TestBreakup(t *testing.T) {
	examples := []struct {
		name string
		s    string
		want []line
	}{
		{
			name: "unbroken string",
			s:    "ABCABBA",
			want: []line{
				{number: 0, text: "A"},
				{number: 1, text: "B"},
				{number: 2, text: "C"},
				{number: 3, text: "A"},
				{number: 4, text: "B"},
				{number: 5, text: "B"},
				{number: 6, text: "A"},
			},
		},
		{
			name: "multiple lines",
			s:    "there once was a man from peru,\nwho dreamed he was eating his shoe,\nhe woke with a fright\nin the middle of the night\nand found that his dream had come true",
			want: []line{
				{number: 0, text: "there once was a man from peru,"},
				{number: 1, text: "who dreamed he was eating his shoe,"},
				{number: 2, text: "he woke with a fright"},
				{number: 3, text: "in the middle of the night"},
				{number: 4, text: "and found that his dream had come true"},
			},
		},
	}
	for _, ex := range examples {
		t.Run(ex.name, func(t *testing.T) {
			got := breakup(ex.s)
			for i, gotLine := range got {
				if gotLine != ex.want[i] {
					t.Fatalf("got %v,\n want %v\nFor line %v\n", gotLine, ex.want[i], i)
				}
			}
		})
	}
}

func TestShortestEdit(t *testing.T) {
	examples := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{
			name: "no change",
			a:    "hello",
			b:    "hello",
			want: 0,
		},
		/*
			{
				name: "one letter",
				a:    "hope",
				b:    "pope",
				want: 1,
			},
			{
				name: "one line",
				a:    "Whitecaps on the bay:\nA broken signboard banging\nIn the April wind.",
				b:    "Whitecaps in the sea:\nA broken signboard banging\nIn the April wind.",
				want: 1,
			},
			{
				name: "cooglan example",
				a:    "ABCABBA",
				b:    "CBABAC",
				want: 5,
			},
		*/
	}
	for _, ex := range examples {
		t.Run(ex.name, func(t *testing.T) {
			got := shortestEdit(ex.a, ex.b)
			if got != ex.want {
				t.Fatalf("got %v, want %v\n", got, ex.want)
			}
		})
	}

}
