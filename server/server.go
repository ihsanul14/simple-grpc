package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	proto "simple-grpc/proto/test"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 30001, "The server port")
)

type TestServer struct {
	proto.UnimplementedTestServiceServer
}

func (s *TestServer) TestGet(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	var err error
	log.Println("testget :", req)
	res := &proto.Response{Id: req.GetId()}
	return res, err
}
func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var srvOpts []grpc.ServerOption

	grpcServer := grpc.NewServer(srvOpts...)
	proto.RegisterTestServiceServer(grpcServer, &TestServer{})

	grpcServer.Serve(lis)
}
