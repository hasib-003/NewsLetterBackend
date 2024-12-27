package repository

import (
	"fmt"
	"github.com/hasib-003/NewsLetterBackend/publisherService/internal/model"
	"gorm.io/gorm"
)

type PublisherRepository struct {
	DB *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) *PublisherRepository {
	return &PublisherRepository{
		DB: db,
	}
}
func (repo *PublisherRepository) CreatePublication(publication *model.Publication) (*model.Publication, error) {
	err := repo.DB.Create(publication).Error
	if err != nil {
		return nil, fmt.Errorf("cannot create publication: %v", err)
	}
	return publication, nil
}
func (repo *PublisherRepository) CreatePost(publicationId uint, post *model.Post) (*model.Post, error) {
	var publication model.Publication
	if err := repo.DB.First(&publication, publicationId).Error; err != nil {
		return nil, fmt.Errorf("cannot Find publication: %v", err)
	}
	post.PublicationId = publicationId
	if err := repo.DB.Create(post).Error; err != nil {
		return nil, fmt.Errorf("cannot create post: %v", err)
	}
	return post, nil
}
