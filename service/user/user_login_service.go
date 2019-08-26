package service

import (
	"blog-admin-service/model"
	"blog-admin-service/serializer"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// token claims
type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

// Login 用户登录函数
func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}
	return user, nil
}

func (service *UserLoginService) GenerateUserToken(user model.User) (string, error) {

	//// 设置token
	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 7200).Unix(),
			Issuer:    "evan zhang",
			Id:        user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	hmacSampleSecret := []byte("newtrekWang")
	return token.SignedString(hmacSampleSecret)

}
