package main

import (
	"device/internal/consumer"
	"device/protos/deviceproto"
	"device/services"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ls, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	server := services.NewService()
	s := grpc.NewServer()
	deviceproto.RegisterDevicesServer(s, server)
	reflection.Register(s)

	fmt.Printf("server started on the port 8081\n")
	a:=consumer.NewSub()
	go func() {
		a.Subscriber()
	}()
	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
