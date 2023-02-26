package main

import (
	"fmt"
	"log"
	"net"

	goclient_proto "goclient/proto"
	rpcfuncs "goclient/rpcfuncs"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World.")

	ss, err := rpcfuncs.NewService()
	if err != nil {
		log.Fatal("[Man] failed to create service: ", err);
	}

	startGrpcServer(&ss)
}

func startGrpcServer(gClient *rpcfuncs.Service) {
	grpcPort := 50096
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Transport [tcp] Listening on :%d", grpcPort)
	grpcServer := grpc.NewServer()
	fmt.Println("[Main] Service object created")

	goclient_proto.RegisterGoClientServer(grpcServer, gClient)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}