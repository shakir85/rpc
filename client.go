package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	args := Args{10, 5}
	var result int

	err = client.Call("Calculator.Add", args, &result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Addition: %d + %d = %d\n", args.A, args.B, result)

	args = Args{20, 8}
	err = client.Call("Calculator.Subtract", args, &result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Subtraction: %d - %d = %d\n", args.A, args.B, result)
}
