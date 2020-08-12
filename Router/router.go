package Router

import (
	"collectbackend/Controllers"
	"collectbackend/Middlewares"
	"collectbackend/Sessions"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	v1 := router.Group("v1")
	{
		v1.POST("/testinsert", Controllers.TestInsert)
		v1.GET("/infopage", Controllers)
	}

	router.Run(":8080")
}
