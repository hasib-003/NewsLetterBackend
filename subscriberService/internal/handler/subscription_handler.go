package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/service"
	"net/http"
)

type SubscriptionInternalHandler struct {
	service *service.SubscriptionService
}

func NewSubscriptionInternalHandler(service *service.SubscriptionService) *SubscriptionInternalHandler {
	return &SubscriptionInternalHandler{
		service: service,
	}
}

func (h *SubscriptionInternalHandler) GetAllSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, "fjf")
}
