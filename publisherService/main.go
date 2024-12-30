package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/common/config"
	"github.com/hasib-003/NewsLetterBackend/publisherService/internal/model"
	"github.com/hasib-003/NewsLetterBackend/publisherService/router"
	"log"
)

func main() {
	config.ConnectDB()
	err := config.DB.AutoMigrate(&model.Publication{}, &model.Post{})
	if err != nil {
		panic(err)
	}
	server := gin.Default()
	router.RegisterRoutes(server, config.GetDB())
	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
	log.Println("server started at port 8080")
}
