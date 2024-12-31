package handler

import (
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/model"
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/service"
)

type SubscriptionInternalHandler struct {
	service *service.SubscriptionService
}

func NewSubscriptionInternalHandler(service *service.SubscriptionService) *SubscriptionInternalHandler {
	return &SubscriptionInternalHandler{
		service: service,
	}
}

func (h *SubscriptionInternalHandler) GetAllSubscription() (*model.Subscription, error) {
	return nil, nil
}
