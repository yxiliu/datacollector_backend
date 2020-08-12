package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndxRecord(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    id,
	})

}

func IndxInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    id,
	})
}
