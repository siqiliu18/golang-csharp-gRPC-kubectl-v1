package rpcmockqhd

import (
	"fmt"
)

// Service is servce
type Service struct {
	Str string
}

// NewService is new service
func NewService() (Service, error) {
	svc := Service {
		Str: "mockqhd",
	}
	fmt.Println("Mockqhd NewService")
	return svc, nil
} 