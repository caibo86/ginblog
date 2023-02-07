package model

import (
	"encoding/base64"
	"github.com/caibo86/ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(128);not null" json:"password"`
	Role     int    `gorm:"type:int;not null" json:"role"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 添加用户
func CreateUser(u *User) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(u).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(perPage, page int) []*User {
	var users []*User
	err := db.Limit(perPage).Offset((page - 1) * perPage).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditUser 更新用户
func EditUser(id int, u *User) int {
	m := make(map[string]interface{})
	m["username"] = u.Username
	m["role"] = u.Role
	err := db.Model(&User{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	u.Password = ScryptPw(u.Password)
	return nil
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 32
	salt := []byte{6, 10, 11, 16, 23, 25, 13, 66}
	hash, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	ret := base64.StdEncoding.EncodeToString(hash)
	return string(ret)
}
