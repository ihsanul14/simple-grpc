package server

import (
	"context"
	"flag"
	"fmt"
	"net"
	proto "simple-grpc/proto/test"
	"simple-grpc/usecase"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	port = flag.Int("port", 30001, "The server port")
)

type TestServer struct {
	logger *logrus.Logger
	db     *gorm.DB
	proto.UnimplementedTestServiceServer
}

type Handler struct {
	Usecase usecase.Usecase
	proto.UnimplementedTestServiceServer
}

func (s Handler) TestGet(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	// var resq Response
	var err error
	// s.logger.Debug("testget :", req)
	// res := &proto.Response{Id: req.GetId()}
	res, err := s.Usecase.ReadData(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	// res := &proto.Response{Id: int32(resq.Id), Name: resq.Nama}
	return res, err
}

func NewGRPCServer(logger *logrus.Logger, usecase usecase.Usecase) {
	var s Handler
	s.Usecase = usecase

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		logger.Fatalf("server.server.NewGRPCServer failed to listen: %v", err)
	}
	var srvOpts []grpc.ServerOption

	grpcServer := grpc.NewServer(srvOpts...)
	proto.RegisterTestServiceServer(grpcServer, s)

	logger.Info("grpc server ready in:", *port)
	grpcServer.Serve(lis)
}
