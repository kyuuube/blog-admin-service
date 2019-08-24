package serializer

import (
	"blog-admin-service/model"
	"github.com/satori/go.uuid"
)

// User 用户序列化器
type User struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"user_name"`
	Nickname  string    `json:"nickname"`
	Status    int8      `json:"status"`
	Avatar    string    `json:"avatar"`
	CreatedAt int64     `json:"created_at"`
	Token     string
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User, token string) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
		Token:     token,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user, ""),
	}
}

// BuildUserLoginResponse 序列化用户响应
func BuildUserLoginResponse(user model.User, token string) UserResponse {
	return UserResponse{
		Data: BuildUser(user, token),
	}
}

// BuildUserLoginResponse 序列化用户响应
func BuildUserInfoResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user, ""),
	}
}
