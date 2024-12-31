package repository

import (
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (repo *UserRepository) CreateUser(user interface{}) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	repo.db = repo.db.Debug()
	return nil
}
func (repo *UserRepository) GetPublisherById(publisherId int) (*models.Publisher, error) {
	var publisher models.Publisher
	if err := repo.db.Where("id = ?", publisherId).First(&publisher).Error; err != nil {
		return nil, err
	}
	return &publisher, nil
}
