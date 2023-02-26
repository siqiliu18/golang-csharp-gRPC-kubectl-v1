package rpcfuncs

import (
	"context"
	"fmt"
	"log"

	pptserver_proto "goclient/ppt-proto"
	goclient_proto "goclient/proto"
)

// SayHello is to say hello
func (s *Service) SayHello(ctx context.Context, in *goclient_proto.Request) (*goclient_proto.Response, error) {
	fmt.Println(">>> SayHello >>>")

	req := pptserver_proto.TemplateRequest{
		Name: "Jasper",
	}

	res, err := s.PptClient.RetrieveTemplate(ctx, &req)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return nil, err
	}

	fmt.Printf("PPT service return: " + res.Message);

	out := goclient_proto.Response{
		Str: "Hello " + in.GetStr() + " | " + res.Message,
	}

	return &out, nil
}