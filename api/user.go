package api

import (
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {

	c.JSON(200, "sss")

}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	c.JSON(200, "绅士手")
}
