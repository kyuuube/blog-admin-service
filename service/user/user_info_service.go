package service

import (
	"blog-admin-service/model"
	"blog-admin-service/serializer"
	"fmt"
)

type UserInfoService struct {
	ID string `form:"id" json:"id" binding:"required,min=5,max=300"`
}

func (service *UserInfoService) GetUserInfo(id string) (model.User, *serializer.Response) {
	//var user model.User
	if user, err := model.GetUser(id); err != nil {
		return user, &serializer.Response{
			Status: 40004,
			Msg:    "没有找到用户",
		}
	} else {
		fmt.Print(user)
		return user, nil
	}

}
