package rpcfuncs

import (
	"fmt"
	ppt_server "goclient/ppt-proto"
	"log"

	"google.golang.org/grpc"
)

// Service is service
type Service struct {
	Str string
	PptClient ppt_server.TemplateMakerClient
}

// NewService is new service
func NewService() (Service, error) {
	svc := Service{
		Str: "any",
	}

	fmt.Println("[Dialing...]")
	conn, err := grpc.Dial("secondgrpcservice-project2:50102", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial server failed: %v", err)
	}
	// defer conn.Close()

	svc.PptClient = ppt_server.NewTemplateMakerClient(conn)

	return svc, nil
}