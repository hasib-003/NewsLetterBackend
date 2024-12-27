package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/publisherService/internal/handler"
	"github.com/hasib-003/NewsLetterBackend/publisherService/internal/repository"
	"github.com/hasib-003/NewsLetterBackend/publisherService/internal/service"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	publisherRepository := repository.NewPublisherRepository(db)
	publisherService := service.NewPublisherService(publisherRepository)
	publisherHandler := handler.NewPublisherHandler(publisherService)

	router.POST("/createPublication", publisherHandler.CreatePublication)
	router.POST("/createPost", publisherHandler.CreatePost)
}
