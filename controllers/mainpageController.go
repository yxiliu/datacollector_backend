package controllers

import (
	"collectbackend/models"
	"collectbackend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IndexPageGET GET请求 主页的所有index
func IndexPageGET(c *gin.Context) {
	var cateID string = c.DefaultQuery("cateid", "0")
	var searchName string = c.DefaultQuery("search", "")
	cateIDInt, _ := strconv.Atoi(cateID)
	var cateIDUInt uint = uint(cateIDInt)
	var indexlist []models.Indx
	if searchName != "" {
		indexlist = models.FindAllByName(searchName)
	} else if cateID != "0" {
		indexlist = models.FindAllByCate(cateIDUInt)
	} else {
		indexlist = models.FindAllIndx()
	}
	c.JSON(http.StatusOK, gin.H{
		"data": indexlist,
	})
}


// NewIndx 一个新的index
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


// GetAllCategory 拿到所有的分类
func GetAllCategory(c *gin.Context) {
	var cat models.Category
	catlist, err := cat.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": catlist})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 300, "data": "", "err": err.Error})
	}
}
