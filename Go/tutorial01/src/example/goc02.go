package main

import (
	"fmt"
	"time"
)

func show(ch chan int, quit chan bool) {
	for {
		select {
		case ret := <-ch:
			fmt.Print(ret, ", ")
			time.Sleep(time.Second * 1)
		case <-quit:
			fmt.Print("\nQuit.")
			break
		}
	}
}

func main() {
	ch := make(chan int, 4)
	quit := make(chan bool)

	go show(ch, quit)

	for i := 0; i < 20; i++ {
		ch <- i
	}

	// block & wait here to let show finish
	quit <- false

}
