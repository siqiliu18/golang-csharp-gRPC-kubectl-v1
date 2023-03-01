package rpcmockqhd

import (
	"context"
	"fmt"

	qhd_proto "mockqhd/proto"
)

// CallQil is to call QIL
func (s *Service) CallQil(ctx context.Context, in *qhd_proto.QILRequest) (*qhd_proto.QILResponse, error) {
	fmt.Println(">>> In mock qhd Call Qil")
	out := qhd_proto.QILResponse{
		Data: "qil response",
	}
	return &out, nil
}
