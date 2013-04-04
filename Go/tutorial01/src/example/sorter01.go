package main

import (
	"fmt"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// declare two types
type Xi []int
type Xs []string

// Xi is implementing Sorter here.
func (p Xi) Len() int {
	return len(p)
}
func (p Xi) Less(i int, j int) bool {
	return p[j] < p[i]
}
func (p Xi) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// Xs is implementing Sorter here.
func (p Xs) Len() int {
	return len(p)
}
func (p Xs) Less(i int, j int) bool {
	return p[j] < p[i]
}
func (p Xs) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// Bubble Sort
func BubbleSort(x Sorter) {
	for i := 0; i < x.Len(); i++ {
		for j := 0; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

func Max(x Sorter) int {
	pos := 0
	for i := 0; i < x.Len(); i++ {
		if x.Less(i, pos) {
			pos = i
		}
	}
	return pos
}

func main() {
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}

	BubbleSort(ints)
	fmt.Printf("%v\n", ints)

	BubbleSort(strings)
	fmt.Printf("%v\n", strings)

	fmt.Printf("Max ints %v\n", ints[Max(ints)])
	fmt.Printf("Max strings %v\n", strings[Max(strings)])
}
