package service

import (
	"blog-admin-service/model"
	"blog-admin-service/serializer"
)

type UserQueryService struct {
	KeyWords  string `form:"keyWords" json:"keyWords" binding:"max=300"`
	Role      int8   `form:"Role" json:"Role"`
	CreatTime int64  `form:"creatTime" json:"creatTime"`
	Limit     int    `form:"limit"`
	Start     int    `form:"start"`
}

func (service *UserQueryService) GetUserList() serializer.Response {
	users := []model.User{}
	total := 0

	if err := model.DB.Where(model.User{CreatedAt: service.CreatTime, Role: service.Role}).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Model(&users).Where("nickname LIKE ?", "%"+service.KeyWords+"%").Order("id desc").Offset(service.Start).Limit(service.Limit).Find(&users).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildUserList(users), uint(total))
}
