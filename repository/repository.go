package repository

import (
	"context"

	pb "simple-grpc/proto/test"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repo struct {
	logger *logrus.Logger
	db     *gorm.DB
}

type Response struct {
	Id         int    `json:"id" gorm:"id"`
	Nama       string `json:"nama" gorm:"nama"`
	Nomor      int    `json:"nomor" gorm:"nomor"`
	Created_at string `json:"created_at" gorm:"created_at"`
	Updated_at string `json:"updated_at" gorm:"updated_at"`
}

type Repository interface {
	GetData(ctx context.Context, request *pb.Request) (res *pb.Response, err error)
}

func NewRepository(logger *logrus.Logger, dbconn *gorm.DB) Repository {
	return &Repo{db: dbconn, logger: logger}
}
func (r Repo) GetData(ctx context.Context, param *pb.Request) (*pb.Response, error) {
	var (
		result []Response
		res    *pb.Response
		err    error
	)
	query := r.db.Table("testing")
	if param.Id != 0 {
		query = query.Where("id = ?", param.Id)
	}
	err = query.Scan(&result).Error
	if err != nil {
		r.logger.Error(err.Error())
		return res, err
	}
	res = &pb.Response{Id: int32(result[0].Id), Name: result[0].Nama}
	return res, err
}
