package main

import (
	"fmt"
	"log"
	"net"
	"user/config"
	"user/internal/connection"
	"user/protos/userproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.Connection()
	ls, err := net.Listen("tcp", fmt.Sprintf(":%s",c.MCPort))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	server := connection.NewServer()
	userproto.RegisterUserServer(s, server)
	reflection.Register(s)

	fmt.Printf("server started on the port 8081")

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
