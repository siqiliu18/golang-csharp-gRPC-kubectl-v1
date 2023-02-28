package main

import (
	"fmt"
	"log"
	"net"

	qhd_proto "mockqhd/proto"
	rpcmockqhd "mockqhd/rpcmockqhd"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("main mock qhd")

	ss, err := rpcmockqhd.NewService()
	if err != nil {
		fmt.Println("[Main] failed to create service: ", err)
	}

	startGrpcServer(&ss)
}

func startGrpcServer(gServer *rpcmockqhd.Service) {
	grpcPort := 50097
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Transport [tcp] Listening on :%d", grpcPort)
	grpcServer := grpc.NewServer()
	fmt.Println("[Main] Service object created")

	qhd_proto.RegisterQHDataServer(grpcServer, gServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal((err))
	}
}