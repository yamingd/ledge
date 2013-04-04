package main

import (
	"example/stack"
	"fmt"
	"unicode/utf8"
)

func fizzBuzz() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%v:%v\n", i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Printf("%v:%v\n", i, "Fizz")
		} else if i%5 == 0 {
			fmt.Printf("%v:%v\n", i, "Buzz")
		}
	}
}

func prymaid() {
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v", "A")
		}
		fmt.Printf("\n")
	}
}

func countLen() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("%v's len = %v, %v\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
}

func avg(sl []float64) float64 {
	result := 0.0
	if len(sl) == 0 {
		return result
	}
	sum := 0.0
	for i := 0; i < len(sl); i++ {
		sum += sl[i]
	}
	result = sum / float64(len(sl))
	return result
}

func deferTest() {
	fmt.Printf("Defer is a Stack\n")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("\t%v", i)
	}
	fmt.Printf("Start:")
}

func stackTest() {
	list := []int{0, 0, 0}
	sl := list[:]
	count := stack.Size(sl)
	fmt.Printf("size = %v\n", count)

	count, sl = stack.Push(sl, 3)
	fmt.Printf("After Push, size = %v\n", count)
	fmt.Printf("Stack is %v\n", stack.String(sl))

	count, sl = stack.Push(sl, 34)
	fmt.Printf("After Push, size = %v\n", count)
	fmt.Printf("Stack is %v\n", stack.String(sl))

	val, sl := stack.Pop(sl)
	fmt.Printf("Pop %v, Size=%v \n", val, stack.Size(sl))

	fmt.Printf("Stack is %v\n", stack.String(sl))
}

func printArgs(args ...int) {
	for index, val := range args {
		fmt.Printf("%v:%v\n", index, val)
	}
}

func fib(list []int, num int) (result []int) {
	count := len(list)
	sum := list[count-2] + list[count-1]
	result = append(list, sum)
	if count+1 <= num {
		result = fib(result, num)
	}
	return
}

func printSlice(list []int) {
	for index, val := range list {
		fmt.Printf("%v = %v\n", index, val)
	}
}

func imap(f func(int) int, list []int) (result []int) {
	for index, val := range list {
		list[index] = f(val)
	}
	result = list
	return
}

func min(list []int) (result int) {
	result = list[0]
	for i := 0; i < len(list); i++ {
		if list[i] < result {
			result = list[i]
		}
	}
	return
}

func max(list []int) (result int) {
	result = list[0]
	for i := 0; i < len(list); i++ {
		if list[i] > result {
			result = list[i]
		}
	}
	return
}

func bubbleSort(index int, list []int) (result []int) {
	ref := list[index]
	done := false
	for i := index + 1; i < len(list); i++ {
		if list[i] < ref {
			list[index] = list[i]
			list[i] = ref
			break
		}
	}
	if index == len(list)-1 {
		done = true
	}
	if !done {
		result = bubbleSort(index+1, list)
	} else {
		result = list
	}
	return
}

func bubblesort2(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}

func addTen(val int) (result int) {
	result = val + 10
	return
}

func fibTest() {
	list := []int{1, 1}
	num := 15
	list = fib(list, num)
	printSlice(list)
}

func imapTest() {
	list := []int{1, 2, 3, 4, 2, 6, 7, 3, 9}
	list = imap(addTen, list)
	printSlice(list)

	fmt.Printf("Min is %v\n", min(list))
	fmt.Printf("Max is %v\n", max(list))
}

func makeFunc(salt int) (f func(int) int) {
	f = func(val int) (ret int) {
		ret = val + salt
		return ret
	}
	return
}

func makeFuncTest() {
	f1 := makeFunc(4)
	fmt.Printf("Add 5 = %v, %v \n", f1(10), f1(11))

	f2 := makeFunc(6)
	fmt.Printf("Add 10 = %v, %v \n", f2(20), f2(32))
}

func bubbleSortTest() {
	list := []int{1, 2, 3, 4, 2, 6, 7, 3, 9}
	println("Bubble Sort Test...")
	bubblesort2(list)
	printSlice(list)
}

func main() {
	// sl := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 1.414}
	// result := avg(sl)
	// fmt.Printf("avg = %v\n", result)

	// deferTest()

	//stackTest()

	//printArgs(1, 2, 3, 4, 5, 6)

	//fibTest()

	//makeFuncTest()

	bubbleSortTest()
}
