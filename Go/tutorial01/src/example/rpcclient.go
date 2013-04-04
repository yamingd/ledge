package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func sycall(num int, client *rpc.Client) {
	args := &Args{num, 8}
	var reply int
	err := client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %vx%v = %v\n", args.A, args.B, reply)
}

func asycall(num int, client *rpc.Client) {
	args := &Args{num, 8}
	quo := new(Quotient)
	divCall := client.Go("Arith.Divide", args, &quo, nil)
	replyCall := <-divCall.Done
	fmt.Printf("%v Arith: %v, %v/%v = %v, %v\n", time.Now(), replyCall.Reply, args.A, args.B, quo.Quo, quo.Rem)
}

func main() {
	ip := "127.0.0.1"
	client, err := rpc.DialHTTP("tcp", ip+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	for i := 0; i < 100; i++ {
		//go sycall(i, client)
		asycall(i, client)
		asycall(i+1000, client)
	}
}
