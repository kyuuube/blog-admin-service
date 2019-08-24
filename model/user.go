package model

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户模型
type User struct {
	ID             uuid.UUID
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         int8
	Avatar         string `gorm:"size:1000"`
	CreatedAt      time.Time
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 15
	// Active 激活用户
	Active int8 = 1
	// Inactive 未激活用户
	Inactive int8 = 2
	// Suspend 被封禁用户
	Suspend int8 = 3
)

// GetUser 用ID获取用户
func GetUser(ID string) (User, error) {
	var user User
	result := DB.Where("id = ?", ID).First(&user)
	fmt.Print(user)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
