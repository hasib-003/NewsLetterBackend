package models

type User struct {
	UserType string `json:"user_type" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Subscriber struct {
	UserType string `json:"user_type" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Type     string `json:"type" binding:"required"`
}

type Publisher struct {
	UserType string `json:"user_type" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Tire     string `json:"tire" binding:"required"`
}

type Admin struct {
	UserType string `json:"user_type" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
