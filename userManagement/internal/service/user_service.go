package service

import (
	"fmt"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/models"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
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
