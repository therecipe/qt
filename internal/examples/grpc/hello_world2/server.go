//profile.proto and this code is taken from: https://github.com/kenshaw/go-jakarta/tree/master/02-gomobile-and-grpc

package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type HelloServer struct{}

// SayHello says 'hi' to the user.
func (hs *HelloServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	// create response
	res := &HelloResponse{
		Reply: fmt.Sprintf("hello %s from go", req.Greeting),
	}

	return res, nil
}

func serve() {
	var err error

	// create socket listener
	l, err := net.Listen("tcp", ":8833")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// create server
	helloServer := &HelloServer{}

	// register server with grpc
	s := grpc.NewServer()
	RegisterHelloServiceServer(s, helloServer)

	// run
	s.Serve(l)
}
