package router

import (
	"collectbackend/controllers"
	middlewares "collectbackend/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())
	// 使用 session(cookie-based)
	// router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	v1 := router.Group("backend")
	{
		// v1.POST("/testinsert", controllers.IndexPage)
		v1.GET("/indexpage", controllers.IndexPage)
		v1.GET("/getallcate", controllers.GetAllCategory)
		// v1.GET("/getdetail",)
	}

	router.Run(":8080")
}
