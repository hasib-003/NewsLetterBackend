package models

import "gorm.io/gorm"

type Subscriber struct {
	gorm.Model
	UserType       string `json:"user_type" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	SubscriberType string `json:"subscriber_type" binding:"required"`
}
type Publisher struct {
	gorm.Model
	UserType        string `json:"user_type" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	Tire            string `json:"tire" binding:"required"`
	SubscriberCount int64  `json:"subscriber_count"`
}
type Admin struct {
	gorm.Model
	UserType string `json:"user_type" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
