package api

import (
	"github.com/gin-gonic/gin"
)

// 毫无用处的接口
func Ping(c *gin.Context) {

	c.JSON(200, "pong")

}
