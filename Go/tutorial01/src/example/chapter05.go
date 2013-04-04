package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func byteRune() {
	fmt.Print("Byte of Yaming: ")
	bs1 := []byte("Yaming")
	for _, val := range bs1 {
		fmt.Printf("%v,", val)
	}
	fmt.Println()

	fmt.Print("Rune of Yaming: ")
	bs2 := []rune("Yaming")
	for _, val := range bs2 {
		fmt.Printf("%v,", val)
	}
	fmt.Println()

	fmt.Print("Byte of 亚明: ")
	bs1 = []byte("亚明")
	for _, val := range bs1 {
		fmt.Printf("%v,", val)
	}
	fmt.Println()

	fmt.Print("Rune of 亚明: ")
	bs2 = []rune("亚明")
	for _, val := range bs2 {
		fmt.Printf("%v,", val)
	}
}

func main() {
	var p *int
	fmt.Printf("p is %v\n", p)

	num := 10
	p = &num
	fmt.Printf("p is %v, value is %v\n", p, *p)

	*p = 20
	fmt.Printf("p is %v, value is %v\n", p, *p)

	person := new(Person)
	person.name = "Yaming"
	person.age = 29

	fmt.Printf("Person is %v\n", person)
	fmt.Printf("Your name is %v\n", person.name)

	//byteRune()

	var p1 Person
	p2 := new(Person)

	fmt.Printf("Person is %v\n", p1)
	fmt.Printf("Person is %v\n", p2)

	p1.name = "P1"
	p2.name = "P2"

	fmt.Printf("Your name is %v\n", p1.name)
	fmt.Printf("Your name is %v\n", p2.name)

	p2 = &p1
	fmt.Printf("Your name is %v\n", p1.name)
	fmt.Printf("Your name is %v\n", p2.name)
}
