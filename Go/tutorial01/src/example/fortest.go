package main

import (
	"fmt"
)

func main() {

	println("For Tutorial.\n")

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
		if i == 50 {
			break
		}
		fmt.Printf("i = %v\t", i)
		if (i+1)%4 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\nSum = %v\n", sum)

	list := []string{"a", "b", "c", "d", "e", "f"}
	for index, value := range list {
		fmt.Printf("index=%v, value=%v\n", index, value)
	}

}
