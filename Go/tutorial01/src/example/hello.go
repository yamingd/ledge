package main

import (
	"example/newmath"
	"fmt"
)

func main() {
	fmt.Printf("Hello, world, Sqrt(2) = %v\n", newmath.Sqrt(2))

	a := 15
	str := "hello world"
	fmt.Printf("a= %v, b=%v\n", a, str)

	var (
		userId   int
		userName string
		isLocked bool
	)

	a, b := 20, 30

	fmt.Printf("a=%v, b=%v, userId=%v, userName=%v, isLocked=%v\n", a, b, userId, userName, isLocked)

	ca := []rune(str)
	ca[0] = 'H'
	str2 := string(ca)

	fmt.Printf("str2=%v, str=%v\n", str2, str)

	str3 := "Start and" +
		" End Here"
	fmt.Printf("str3 = %v\n", str3)

	str4 := `Start and
			 End Here`
	fmt.Printf("str4 = %v\n", str4)

	var cnumber complex64
	cnumber = 5 + 5i
	fmt.Printf("cnumber=%v\n", cnumber)

	if true && true {
		fmt.Printf("We are here. True")
	}

	if !false {
		fmt.Printf("We are still here. False\n")
	}

}
