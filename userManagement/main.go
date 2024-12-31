package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/common/config"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/models"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	config.ConnectDB()
	subscriberConn, err := grpc.NewClient("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func(subscriberConn *grpc.ClientConn) {
		err := subscriberConn.Close()
		if err != nil {

		}
	}(subscriberConn)

	err = config.DB.AutoMigrate(&models.Publisher{}, &models.Admin{}, &models.Subscriber{})
	if err != nil {
		panic(err)
	}
	server := gin.Default()
	router.RegisterRoutes(server)
	log.Println("server start......")
	err = server.Run(":8081")
	if err != nil {
		panic(err)
	}

}
