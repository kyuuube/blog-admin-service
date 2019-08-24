package api

import (
	"blog-admin-service/serializer"
	service "blog-admin-service/service/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// UserRegister
// @Summary 用户注册接口
// @Description register account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param body body service.UserRegisterService true "register account"
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin
// @Summary 用户登录接口
// @Description get accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param body body service.UserLoginService true "login account"
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Login(); err != nil {
			c.JSON(200, err)
		} else {
			//// 设置token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"foo": "bar",
				"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
			})

			hmacSampleSecret := []byte("newtrekWang")
			tokenString, err := token.SignedString(hmacSampleSecret)
			if err != nil {
				c.JSON(200, err)
			}
			res := serializer.BuildUserLoginResponse(user, tokenString)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// VenaUsor
// @Summary 获取用户信息
// @Description get current uer info
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id query string false "用户id"
// @Router /api/v1/user/me [get]
func VenaUsor(c *gin.Context) {
	var service service.UserInfoService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.GetUserInfo(service.ID); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildUserInfoResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
