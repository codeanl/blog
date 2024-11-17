package model

import (
	"errors"

	"gorm.io/gorm"
)

// TODO TDM还不知道有哪些字段 先这样吧
type User struct {
	Model
	Email    string `json:"email" gorm:"type:varchar(30)"`
	Nickname string `json:"nickname" gorm:"type:varchar(30);not null"`
	Username string `json:"username" gorm:"unique;type:varchar(30);not null"`
	Password string `json:"-" gorm:"type:varchar(100);not null"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);not null"`
}

func CheckEmailExist(db *gorm.DB, email string) (bool, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func GetUserInfoByUsername(db *gorm.DB, name string) (*User, error) {
	var user User
	result := db.Where("username = ?", name).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, result.Error
}

func Register(db *gorm.DB, user *User) error {
	result := db.Model(&user).Create(&user)
	return result.Error
}
