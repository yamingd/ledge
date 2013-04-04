/*
  Stack implements a LIFO collections
  has Size, Push, Pop and String four functions.
*/
package stack

import (
	"fmt"
)

// 
// return the total elements of a stack.
//
func Size(list []int) (count int) {
	count = len(list)
	return
}

//
// push a new element into stack.
//
func Push(list []int, val int) (count int, result []int) {
	result = append(list, val)
	count = len(result)
	return
}

func Pop(list []int) (val int, result []int) {
	index := len(list) - 1
	val = list[index]
	result = list[0:index]
	return
}

func String(list []int) (result string) {
	for i := 0; i < len(list); i++ {
		result += fmt.Sprintf("[%v:%v]", i, list[i])
	}
	return
}
