package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/NewsLetterBackend/publisherService/internal/model"
	"github.com/hasib-003/NewsLetterBackend/publisherService/internal/service"
	"net/http"
	"strconv"
)

type PublisherHandler struct {
	PublisherService *service.PublisherService
}

func NewPublisherHandler(publisherService *service.PublisherService) *PublisherHandler {
	return &PublisherHandler{
		PublisherService: publisherService,
	}
}
func (handler *PublisherHandler) CreatePublication(c *gin.Context) {
	var publication model.Publication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publicationlist, err := handler.PublisherService.CreatePublication(&publication)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": publicationlist})

}
func (handler *PublisherHandler) CreatePost(c *gin.Context) {
	publicationId, err := strconv.Atoi(c.Param("publicationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postlist, err := handler.PublisherService.CreatePost(uint(publicationId), &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": postlist})
}
