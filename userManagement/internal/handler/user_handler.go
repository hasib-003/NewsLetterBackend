package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/models"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/service"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
func (handler *UserHandler) CreateUser(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while binding data"})
		return
	}
	switch userInput.UserType {
	case "admin":
		var admin models.Admin
		admin = models.Admin{
			UserType: userInput.UserType,
			Email:    userInput.Email,
			Password: userInput.Password,
		}
		if err := c.ShouldBindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while binding admin data"})
			return
		}
		if err := handler.userService.CreateUser("admin", &admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while create admin user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "admin created successfully"})
	case "subscriber":
		var subscriber models.Subscriber
		if err := c.ShouldBindJSON(&subscriber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while binding subscriber data"})
			return
		}
		if err := handler.userService.CreateUser("subscriber", &subscriber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while create subscriber user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "subscriber created successfully"})
	case "publisher":
		var publisher models.Publisher
		if err := c.ShouldBindJSON(&publisher); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while binding publisher data"})
			return
		}
		if err := handler.userService.CreateUser("publisher", &publisher); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while create publisher user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "publisher created successfully"})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user type"})
	}

}
