package client

import (
	"testing"
)

func TestGrpcService_GrpcWrite(t *testing.T) {
	GrpcService{}.GrpcWrite(":9002")
}
