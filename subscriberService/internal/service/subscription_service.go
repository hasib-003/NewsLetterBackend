package service

import "github.com/hasib-003/NewsLetterBackend/subscriberService/internal/repository"

type SubscriptionService struct {
	repository *repository.SubscriptionRepository
}

func NewSubscriptionService(repository *repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		repository: repository,
	}
}

func (s *SubscriptionService) SubscribeToPublication(userId int64, publicationId int64) error {
	err := s.repository.SubscribeToPublication(userId, publicationId)
	if err != nil {
		return err
	}
	return nil
}
func (s *SubscriptionService) UnSubscribeToPublication(userId int64, publicationId int64) error {
	err := s.repository.UnSubscribeToPublication(userId, publicationId)
	if err != nil {
		return err
	}
	return nil
}
