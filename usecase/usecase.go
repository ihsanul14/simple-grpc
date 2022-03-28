package usecase

import (
	"context"
	"log"

	pb "simple-grpc/proto/test"
	repo "simple-grpc/repository"

	"github.com/sirupsen/logrus"
)

type UsecaseModul struct {
	logger *logrus.Logger
	Repo   repo.Repository
}

type Usecase interface {
	ReadData(ctx context.Context, request *pb.Request) (res *pb.Response, err error)
}

func NewUsecase(log *logrus.Logger, u repo.Repository) Usecase {
	return &UsecaseModul{logger: log, Repo: u}
}
func (u UsecaseModul) ReadData(ctx context.Context, param *pb.Request) (*pb.Response, error) {
	res, err := u.Repo.GetData(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return res, err
	}
	return res, err
}
