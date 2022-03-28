package delivery

import (
	srv "simple-grpc/server"
	"simple-grpc/usecase"

	"github.com/sirupsen/logrus"
)

func StartApp(logger *logrus.Logger, u usecase.Usecase) {
	srv.NewGRPCServer(logger, u)
}
