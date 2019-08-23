package server

import (
	"blog-admin-service/api"
	_ "blog-admin-service/docs"
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

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)
		//
		//// 视频操作
		//v1.POST("videos", api.CreateVideo)
		//v1.GET("video/:id", api.ShowVideo)
		//v1.GET("videos", api.ListVideo)
		//v1.PUT("video/:id", api.UpdateVideo)
		//v1.DELETE("video/:id", api.DeleteVideo)
		//// 排行榜
		//v1.GET("rank/daily", api.DailyRank)
		//// 其他
		//v1.POST("upload/token", api.UploadToken)
	}
	return r
}
