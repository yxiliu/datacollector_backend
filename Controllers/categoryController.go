package Controllers

import (
	"collectbackend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {

	var cat Models.Category
	cat.FindAll()
	err := c.ShouldBindJSON(&testService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := testService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Insert() error!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    id,
	})

}
