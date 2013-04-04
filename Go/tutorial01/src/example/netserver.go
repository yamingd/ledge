package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8053")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go Echo(c)
		fmt.Println("Wait...")
	}
}

func Echo(c net.Conn) {
	fmt.Println("\tConnected ... ")
	defer c.Close()
	debug := false
	r := bufio.NewReader(c)
	for {
		buf, err := r.ReadBytes('\n')
		if buf[0] == '\n' {
			break
		}
		nr := len(buf)
		if err != nil {
			fmt.Printf("Failed to read. %v\n", err.Error())
			return
		}

		fmt.Printf("\tREV %v\n", string(buf[0:nr-1]))

		nw, err := c.Write(buf)
		if err != nil {
			fmt.Printf("Failed to write. %v\n", err.Error())
			return
		}
		if debug {
			fmt.Printf("Stat r %v, w %v\n", nr, nw)
		}
	}

	fmt.Println("\tFinished ... ")
}
