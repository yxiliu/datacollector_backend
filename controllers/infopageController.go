package controllers

import (
	"collectbackend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IndxRecord 显示一个index的所有数据
func IndxRecordTable(c *gin.Context) {
	var id string = c.DefaultQuery("id","0")
	idint, _ = strconv.Atoi(id)

	c.ShouldBindJSON(&indxrecord)
	indexlist := indxrecord.GetTable()
	c.JSON(http.StatusOK, gin.H{
		"data": indexlist,
	})
}

// IndxInfo 显示一个index的内容信息
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
