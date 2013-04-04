package main

import ()

func output(sl []int) {
	print(sl, " ==> ")
	for i := 0; i < len(sl); i++ {
		print(sl[i], ",")
	}
	println("")
}

func main() {
	sl := make([]int, 10)
	println("len=", len(sl))

	sl[0] = 1
	println("sl[0]=", sl[0])

	s0 := []int{0, 0}
	s1 := append(s0, 1)
	s2 := append(s1, 2, 3, 4)
	s3 := append(s2, s0...)

	output(s0)
	output(s1)
	output(s2)
	output(s3)

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl2 := make([]int, 6)
	n1 := copy(sl2, a[0:])
	print("n1 = ", n1, "\n")
	output(sl2)

	n2 := copy(sl2, sl2[2:])
	print("n2 = ", n2, "\n")
	output(sl2)

	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	year := 0
	for _, days := range monthdays {
		year += days
	}
	println("Number of days in a year:", year)

	map02 := make(map[string]string)
	map02["A"] = "A"
	map02["B"] = "B"
	println("map02[A] => ", map02["A"])

	delete(map02, "A")
	delete(map02, "C")
	println("map02[A] => ", map02["A"])
}
