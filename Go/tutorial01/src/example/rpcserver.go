package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	time.Sleep(2 * time.Second)
	*reply = args.A * args.B
	fmt.Printf("args is %v\n", args)
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	time.Sleep(2 * time.Second)
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	fmt.Printf("%v args is %v\n", time.Now(), args)
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	artith := new(Arith)
	rpc.Register(artith)
	rpc.HandleHTTP()
	li, e := net.Listen("tcp", "127.0.0.1:1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(li, nil)

	for {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
}
