package handlers

import (
	"github.com/gin-gonic/gin"
	"integrations_apis/internal/application/sera4/dto"
	"integrations_apis/internal/application/sera4/service"
	"net/http"
)

type Sera4 struct {
	Service *service.Sera4Service
}

func NewSera4Handler(service *service.Sera4Service) *Sera4 {
	return &Sera4{
		Service: service,
	}
}

func (s *Sera4) CreateUser(c *gin.Context) {
	var user dto.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := s.Service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"userId": userId})
}

func (s *Sera4) GetUser(c *gin.Context) {
	id := c.Param("id")
	data, err := s.Service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data["data"]})
}

func (s *Sera4) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := s.Service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
		"message": "User deleted",
	})
}

func (s *Sera4) CreateKey(c *gin.Context) {
	var key dto.KeyDTO
	if err := c.ShouldBindJSON(&key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	keyId, err := s.Service.CreateKey(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"keyId": keyId})
}

func (s *Sera4) DeleteKey(c *gin.Context) {
	id := c.Param("id")
	err := s.Service.DeleteKey(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Key deleted"})
}
