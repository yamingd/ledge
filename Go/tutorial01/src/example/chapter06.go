package main

import (
	"fmt"
)

type S struct {
	i int
}

func (s *S) Get() int {
	return s.i
}

func (s *S) Put(val int) {
	s.i = val
}

type R struct {
	i int
}

func (s *R) Get() int {
	return s.i
}

func (s *R) Put(val int) {
	s.i = val
}

type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	switch t := p.(type) {
	case *S:
		fmt.Printf("p is *S, %v\n", t)
	case *R:
		fmt.Printf("p is *R, %v\n", t)
	default:
		fmt.Printf("Unknown")
	}

	fmt.Printf("p is %v\n", p.Get())
	p.Put(10)
	fmt.Printf("p is %v\n", p.Get())
}

func main() {
	var s S
	f(&s)

	var r R
	f(&r)

	p := new(S)
	t, ok := p.(I)
	if ok {
		fmt.Printf("S is I, %v\n", t)
	}
}
