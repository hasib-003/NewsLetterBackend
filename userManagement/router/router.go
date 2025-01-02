package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/common/config"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/handler"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/repository"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/service"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/proto/subscription"
	"google.golang.org/grpc"
)

func NewRouter(subscriberConn *grpc.ClientConn) *gin.Engine {

	router := gin.Default()

	db := config.GetDB()
	repo := repository.NewUserRepository(db)
	subscriptionClient := subscription.NewSubscriptionClient(subscriberConn)
	services := service.NewUserService(repo, subscriptionClient)
	handlers := handler.NewUserHandler(services)
	router.POST("/createUser", handlers.CreateUser)

	return router
}
