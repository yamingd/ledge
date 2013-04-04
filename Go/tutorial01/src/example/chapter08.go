package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func example01() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(pwd)

	buf := make([]byte, 1024)
	fname := fmt.Sprintf("%v\\example\\data.log", pwd)
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Can't Open data.log")
		return
	}
	defer f.Close()

	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func example02() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Current Path: %v", pwd)

	buf := make([]byte, 1024)
	fname := fmt.Sprintf("%v\\example\\data.log", pwd)
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Can't Open data.log")
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	// flush data to output, otherwise nothing will be there.
	defer w.Flush()

	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[0:n])
	}
}

func example03() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Current Path: %v\n", pwd)

	fname := fmt.Sprintf("%v\\example\\data.log", pwd)
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Can't Open data.log")
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	// flush data to output, otherwise nothing will be there.
	defer w.Flush()

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			break
		}
		w.WriteString(line)
	}
}

func main() {
	example03()
}
