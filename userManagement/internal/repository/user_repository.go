package repository

import "gorm.io/gorm"

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
