package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/common/config"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/models"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/router"
	"log"
)

func main() {
	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.User{}, &models.Publisher{}, &models.Admin{}, &models.Subscriber{})
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
