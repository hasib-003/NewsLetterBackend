package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/common/config"
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/model"
	"github.com/hasib-003/NewsLetterBackend/subscriberService/routes"
	"google.golang.org/grpc"
	"net"
)

func main() {
	config.ConnectDB()
	err := config.DB.AutoMigrate(&model.Subscription{})
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	routes.RegisterRoutes(router)
	go func() {
		if err := router.Run(":8090"); err != nil {
			panic(err)
		}
	}()
	listener, err := net.Listen("tcp", ":50505")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
