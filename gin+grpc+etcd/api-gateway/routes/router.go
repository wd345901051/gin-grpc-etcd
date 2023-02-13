package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/middlerware"

	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middlerware.MiddlerFunc(), middlerware.InitMiddleware(service))
	v1 := ginRouter.Group("/api/vi")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "seccess")
		})
		// 用户服务
		v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", handler.UserLogin)
		authed := v1.Group("/")
		authed.Use(middlerware.JWT())
		{
			// 任务模块
			authed.GET("task", handler.ListTask)
			authed.POST("task", handler.CreateTask)
			authed.PUT("task", handler.UpdataTask)
			authed.DELETE("task", handler.DeleteTask)
		}
	}
	return ginRouter
}
