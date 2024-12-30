package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/common/config"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/handler"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/repository"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/service"
)

func RegisterRoutes(router *gin.Engine) {
	db := config.GetDB()
	repo := repository.NewUserRepository(db)
	services := service.NewUserService(repo)
	handlers := handler.NewUserHandler(services)

	router.POST("createUser", handlers.CreateUser)
}
