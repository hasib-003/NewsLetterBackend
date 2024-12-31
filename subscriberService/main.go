package main

import (
	"github.com/hasib-003/NewsLetterBackend/common/config"
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/model"
	"google.golang.org/grpc"
)

func main() {
	config.ConnectDB()
	err := config.DB.AutoMigrate(&model.Subscription{})
	if err != nil {
		panic(err)
	}
	grpcserver := grpc.NewServer()

}
