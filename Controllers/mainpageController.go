package Controllers

import (
	"collectbackend/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	var prefaceinterface Services.Preface
	err := c.ShouldBindJSON(&prefaceinterface)
	listOfPreface := prefaceinterface.SearchAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": listOfPreface,
		})
		return
	}
}
