package router

import (
	"collectbackend/controllers"
	middlewares "collectbackend/middlewares"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由初始化
func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())
	// 使用 session(cookie-based)
	// router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	backend := router.Group("backend")
	{
		// v1.POST("/testinsert", controllers.IndexPage)
		backend.GET("/indexpage", controllers.IndexPageGET)
		backend.POST("/searchindex", controllers.IndexPageSearch)
		backend.GET("/getallcate", controllers.GetAllCategory)
		// v1.GET("/getdetail",)
		backend.POST("/newIndex", controllers.NewIndx)

		backend.GET("/getIndexTable", controllers.IndxRecordTable)

		backend.GET("/getIndexInfo", controllers.IndxInfo)
		backend.POST("/newrecord", controllers.NewRecord)
		// backend.GET("/test", controllers.NewCate)
	}

	router.Run(":8080")
}
