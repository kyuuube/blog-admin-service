package server

import (
	"blog-admin-service/api"
	_ "blog-admin-service/docs"
	"blog-admin-service/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		authed := r.Group("/api/v1")
		authed.Use(middleware.JWTAuth())
		{
			authed.GET("user/me", api.VenaUsor)
			authed.GET("user/list", api.UserList)
		}

	}
	return r
}
