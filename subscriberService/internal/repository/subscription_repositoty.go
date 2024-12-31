package repository

import (
	"github.com/hasib-003/NewsLetterBackend/subscriberService/internal/model"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	DB *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		DB: db,
	}
}

func (repo *SubscriptionRepository) SubscribeToPublication(userId int64, publicationId int64) error {
	subscription := &model.Subscription{
		UserId:        userId,
		PublicationId: publicationId,
	}
	err := repo.DB.Create(subscription).Error
	if err != nil {
		return err
	}
	return nil
}
func (repo *SubscriptionRepository) UnSubscribeToPublication(userId int64, publicationId int64) error {
	subscription := &model.Subscription{}
	err := repo.DB.Where("user_id = ? AND publication_id = ?", userId, publicationId).Delete(subscription).Error
	if err != nil {
		return err
	}
	return nil
}
