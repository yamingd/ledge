package main

import (
	"fmt"
	"time"
)

// declare a channel
var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf("%v is ready\n", w)
	// send value to channel
	c <- sec
}

func recvalue() {
	fmt.Println("I'm waiting but not too long.")

	// receive value from channel
	result := <-c
	fmt.Println(result)

	result = <-c
	fmt.Println(result)
}

func startc() {
	go ready("Tea", 5)
	go ready("Dinner", 10)
}

func selectvalue() {
	i := 0
L:
	for {
		select {
		case <-c:
			i++
			if i > 1 {
				break L
			}
		}
	}
}

func main() {
	// init a channel
	c = make(chan int)
	startc()
	//recvalue()
	selectvalue()
}
