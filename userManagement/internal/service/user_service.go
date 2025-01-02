package service

import (
	"context"
	"fmt"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/models"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/repository"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/proto/subscription"
	"log"
)

type UserService struct {
	userRepository     *repository.UserRepository
	subscriptionClient subscription.SubscriptionClient
}

func NewUserService(userRepository *repository.UserRepository, subscriptionClient subscription.SubscriptionClient) *UserService {
	return &UserService{
		userRepository:     userRepository,
		subscriptionClient: subscriptionClient,
	}
}

func (s *UserService) CreateUser(userType string, userInput interface{}) error {
	var user interface{}
	switch userType {
	case "admin":
		user = userInput.(*models.Admin)
	case "subscriber":
		user = userInput.(*models.Subscriber)
	case "publisher":
		user = userInput.(*models.Publisher)
	}
	err := s.userRepository.CreateUser(user)
	if err != nil {
		return fmt.Errorf("create user error: %v", err)
	}
	return nil
}
func (s *UserService) SubscribeToPublication(userId int64, publicationId int64) error {
	//TODO   extract publisherid from token

	publisher, err := s.userRepository.GetPublisherById(1)
	publisher.SubscriberCount -= 1
	req := &subscription.SubscriptionRequest{
		UserId:        userId,
		PublicationId: publicationId,
	}
	_, err = s.subscriptionClient.SubscribeToPublication(context.Background(), req)
	if err != nil {
		return fmt.Errorf("subscribe to publication error: %v", err)
	}
	return nil
}
func (s *UserService) UnsubscribeFromPublication(userId int64, publicationId int64) error {
	// TODO: Extract publisherId from token if required

	publisher, err := s.userRepository.GetPublisherById(1)
	if err != nil {
		return fmt.Errorf("error retrieving publisher: %v", err)
	}

	publisher.SubscriberCount += 1

	req := &subscription.SubscriptionRequest{
		UserId:        userId,
		PublicationId: publicationId,
	}

	_, err = s.subscriptionClient.UnSubscribeToPublication(context.Background(), req)
	if err != nil {
		return fmt.Errorf("unsubscribe from publication error: %v", err)
	}
	return nil
}

func (s *UserService) ListSubscriptions(userId int64) ([]*models.Subscriber, error) {
	req := &subscription.GetSubscriptionsRequest{
		UserId: userId,
	}

	resp, err := s.subscriptionClient.GetAllSubscriptions(context.Background(), req)
	if err != nil {
		log.Println("Error fetching subscriptions:", err)
		return nil, err
	}

	return resp.GetAllSubscriptions(), nil
}
