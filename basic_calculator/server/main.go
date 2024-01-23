package main

import (
	"log"
	pb "basic_calculator/compiled_protos/protos"

	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8081"
)

// grpc server object to reference server generated by grpc plugin for Protocol buffer compiler
// server will be used to implement rpc endpoints
type CalculatorServiceServer struct {
	pb.UnimplementedCalculatorServiceServer
}

func main(){
	listener, err := net.Listen("tcp", port) // function to begin listening the port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()                                   // creating 'server object' from grpc module
	pb.RegisterCalculatorServiceServer(s, &CalculatorServiceServer{})  // register server as new grpc service to implement rpc endpoints
	reflection.Register(s)                                  // for serializing & de-serializing data

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to load server: %v", err)
	}
}