package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/models"
	"github.com/hasib-003/NewsLetterBackend/usermanagement/internal/service"
	"io"
	"net/http"
	"strconv"
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
	// Read the raw body first
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading request body"})
		return
	}

	// Get the user type from the request
	type userTypeWrapper struct {
		UserType string `json:"user_type" binding:"required"`
	}

	var wrapper userTypeWrapper
	if err := json.Unmarshal(bodyBytes, &wrapper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while parsing user type"})
		return
	}

	switch wrapper.UserType {
	case "admin":
		var admin models.Admin
		if err := json.Unmarshal(bodyBytes, &admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while binding admin data"})
			return
		}

		if err := handler.userService.CreateUser("admin", &admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating admin user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "admin created successfully"})

	case "subscriber":
		var subscriber models.Subscriber
		if err := json.Unmarshal(bodyBytes, &subscriber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while binding subscriber data"})
			return
		}

		if err := handler.userService.CreateUser("subscriber", &subscriber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating subscriber user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "subscriber created successfully"})

	case "publisher":
		var publisher models.Publisher
		if err := json.Unmarshal(bodyBytes, &publisher); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while binding publisher data"})
			return
		}

		if err := handler.userService.CreateUser("publisher", &publisher); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating publisher user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "publisher created successfully"})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user type"})
	}
}

func (handler *UserHandler) SubscribeToPublication(c *gin.Context) {
	userIdStr := c.Param("userId")
	publicationIdStr := c.Param("publicationId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while parsing user id"})
	}
	publicationId, err := strconv.ParseInt(publicationIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while parsing publication id"})
	}
	err = handler.userService.SubscribeToPublication(userId, publicationId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (handler *UserHandler) UnsubscribeFromPublication(c *gin.Context) {
	userIdStr := c.Param("userId")
	publicationIdStr := c.Param("publicationId")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while parsing user ID"})
		return
	}

	publicationId, err := strconv.ParseInt(publicationIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while parsing publication ID"})
		return
	}

	err = handler.userService.UnsubscribeFromPublication(userId, publicationId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "unsubscribed from publication successfully"})
}
