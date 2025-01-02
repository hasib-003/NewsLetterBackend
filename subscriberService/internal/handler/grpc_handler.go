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

func (h *SubscriptionHandler) GetAllSubscriptions(ctx context.Context, req *subscription.SubscriptionRequest) (*subscription.GetSubscriptionsResponse, error) {
	subscriptions, err := h.SubscriptionService.GetAllSubscriptions(req.UserId)
	if err != nil {
		return nil, err
	}

	var subscriptionList []*subscription.SubscribedTopic
	for _, sub := range subscriptions {
		subscriptionList = append(subscriptionList, &subscription.SubscribedTopic{
			UserId:        sub,
			PublicationId: 4,
		})
	}

	return &subscription.GetSubscriptionsResponse{
		Subscriptions: subscriptionList,
	}, nil
}
