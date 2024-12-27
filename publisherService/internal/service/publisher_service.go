package service

import (
	"fmt"
	"github.com/hasib-003/NewsLetterBackend/publisherService/internal/model"
	"github.com/hasib-003/NewsLetterBackend/publisherService/internal/repository"
)

type PublisherService struct {
	repository *repository.PublisherRepository
}

func NewPublisherService(repository *repository.PublisherRepository) *PublisherService {
	return &PublisherService{
		repository: repository,
	}
}
func (s *PublisherService) CreatePublication(publication *model.Publication) (*model.Publication, error) {
	publication, err := s.repository.CreatePublication(publication)
	if err != nil {
		return nil, fmt.Errorf("create publication: %v", err)
	}
	return publication, nil
}
func (s *PublisherService) CreatePost(publicationId uint, post *model.Post) (*model.Post, error) {
	post, err := s.repository.CreatePost(publicationId, post)
	if err != nil {
		return nil, fmt.Errorf("create post: %v", err)
	}
	return post, nil
}
