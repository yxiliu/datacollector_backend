package Controllers

import (
	"collectbackend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var cat Models.Category
	catlist := cat.FindAll()
	c.JSON(http.StatusOK, gin.H{"data": catlist})
}
