package main

import (
	"simple-grpc/database"
	"simple-grpc/delivery"
	"simple-grpc/repository"
	"simple-grpc/usecase"

	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	logger := logrus.New()
	Db, err := database.ConnectMySQL(logger)
	if err != nil {
		logger.Fatal("main.database.error :", err)
		return
	}
	repository := repository.NewRepository(logger, Db)
	usecase := usecase.NewUsecase(logger, repository)
	delivery.StartApp(logger, usecase)
}
