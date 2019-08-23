package api

import (
	"github.com/gin-gonic/gin"
)

// UserRegister
// @Summary 用户注册接口
// @Description get accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q" Format(email)
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {

	c.JSON(200, "sss")

}

// UserLogin
// @Summary 用户登录接口
// @Description get accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account query string false "用户账号"
// @Param password query string false "用户密码"
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	c.JSON(200, "绅士手")
}
