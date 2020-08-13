package controllers

import (
	"collectbackend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var cat models.Category
	catlist, _ := cat.FindAll()
	c.JSON(http.StatusOK, gin.H{"data": catlist})
}
