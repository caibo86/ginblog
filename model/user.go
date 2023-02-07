package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int;not null" json:"role"`
}

// 添加用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {

	}
}
