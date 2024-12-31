package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/common/config"
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/handler"
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/repository"
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/service"
)

func RegisterRoutes(router *gin.Engine) {
	db := config.GetDB()
	subscriptionRepository := repository.NewSubscriptionRepository(db)
	subscriptionService := service.NewSubscriptionService(subscriptionRepository)
	subscriptionHandler := handler.NewSubscriptionInternalHandler(subscriptionService)

	router.POST("/getAllSubscriotion", subscriptionHandler.GetAllSubscription)

}
