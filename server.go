package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calculator struct{}

type Args struct {
	A, B int
}

func (c *Calculator) Add(args Args, result *int) error {
	*result = args.A + args.B
	return nil
}

func (c *Calculator) Subtract(args Args, result *int) error {
	*result = args.A - args.B
	return nil
}

func main() {
	calculator := new(Calculator)
	rpc.Register(calculator)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server listening on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
