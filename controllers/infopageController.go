package controllers

import (
	"collectbackend/models"
	"collectbackend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IndxRecord 显示一个index的所有数据
func IndxRecordTable(c *gin.Context) {
	var id string = c.DefaultQuery("id", "0")
	var year string = c.DefaultQuery("year", "2020")
	idint, _ = strconv.Atoi(id)
	yearint, _ = strconv.Atoi(year)
	listofindex := models.FindAllRecordByIndexAndYear(uint(idint), uint(yearint))
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"data": indexlist,
	})
}

// IndxInfo 显示一个index的内容信息
func IndxInfo(c *gin.Context) {
	var id string = c.DefaultQuery("id", "0")

	if id == "0" {
		c.JSON(http.StatusOK, gin.H{
			"status": 3000,
			"data":   "no id",
		})
	} else {
		idint, _ = strconv.Atoi(id)
		var idx models.Indx
		idx.FindOne(uint(idint))
		c.JSON(http.StatusOK, gin.H{
			"status": 2000,
			"data":   idx,
		})
	}
}

// NewRecord 添加一个信息
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
