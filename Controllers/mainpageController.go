package controllers

import (
	"collectbackend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	var prefaceinterface services.Preface
	err := c.ShouldBindJSON(&prefaceinterface)
	listOfPreface := prefaceinterface.SearchAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": listOfPreface,
		})
		return
	}
}

func NewIndx(c *gin.Context) {
	var newthing services.NewIndex
	c.ShouldBindJSON(&newthing)
	id, err := newthing.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": id,
		})
	}
}
