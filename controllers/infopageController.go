package controllers

import (
	"collectbackend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndxRecord(c *gin.Context) {
	var indxrecord services.IndxRecordInfo
	c.ShouldBindJSON(&indxrecord)
	indexlist := indxrecord.GetTable()
	c.JSON(http.StatusOK, gin.H{
		"data": indexlist,
	})

}

func IndxInfo(c *gin.Context) {
	var indxinfo services.IndxInfo
	c.ShouldBindJSON(&indxinfo)
	res := indxinfo.GetEachInfo()
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func NewRecord(c *gin.Context) {
	var newlist []services.NewRecord

	err := c.ShouldBindJSON(&newlist)
	if err != nil {
		for _, x := range newlist {
			x.Insert()
		}
		c.JSON(http.StatusOK, gin.H{
			"data": "ok",
		})
	}

}
