package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	line := "Hello"
	for i := 0; i < 10000; i++ {
		send(fmt.Sprintf("%v %v", line, i))
	}
}

func robot() {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if line == "EOF" {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println("ERROR")
			break
		}
		for i := 0; i < 1000; i++ {
			send(fmt.Sprintf("%v %v", line, i))
		}
	}
}

func send(line string) {
	c, err := net.Dial("tcp", "127.0.0.1:8053")
	if err != nil {
		fmt.Printf("\nFailed to connect. %v", err.Error())
		return
	}
	data := fmt.Sprintf("%v\n\n", line)
	size, err := c.Write([]byte(data))
	if err != nil {
		fmt.Printf("\nFailed to write. %v", err.Error())
		return
	}
	fmt.Printf("\nSEND: %v", size)

	r := bufio.NewReader(c)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("\nFailed to read. %v", err.Error())
			break
		}
		nr := len(buf)
		fmt.Printf("\nEcho >> %v(%v)", string(buf[0:nr-1]), nr)
		if nr == 0 {
			break
		}
	}

	fmt.Println("\nFinished ... ")
}
