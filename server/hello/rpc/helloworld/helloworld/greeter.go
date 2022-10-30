package main

import (
	"context"

	pb "git.code.oa.com/gatesma/demo-go/hello/rpc/hello"
)

type greeterImpl struct{}

func (s *greeterImpl) SayHello(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
	// implement business logic here ...
	// ...
	return nil
}

func (s *greeterImpl) SayHi(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
	// implement business logic here ...
	// ...
	return nil
}
