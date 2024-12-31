package handler

import (
	"context"
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/service"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/proto/subscription"
)

type SubscriptionHandler struct {
	subscription.UnimplementedSubscriptionServer
	SubscriptionService *service.SubscriptionService
}

func (h *SubscriptionHandler) SubscribeToPublication(ctx context.Context, req *subscription.SubscriptionRequest) (*subscription.SubscriptionResponse, error) {
	err := h.SubscriptionService.SubscribeToPublication(req.UserId, req.PublicationId)
	if err != nil {
		return nil, err
	}
	return &subscription.SubscriptionResponse{
		Success: true,
		Message: "successfully subscribed to publication",
	}, nil
}
func (h *SubscriptionHandler) UnSubscribeToPublication(ctx context.Context, req *subscription.SubscriptionRequest) (*subscription.SubscriptionResponse, error) {
	err := h.SubscriptionService.UnSubscribeToPublication(req.UserId, req.PublicationId)
	if err != nil {
		return nil, err
	}
	return &subscription.SubscriptionResponse{
		Success: true,
		Message: "successfully subscribed to publication",
	}, nil
}
