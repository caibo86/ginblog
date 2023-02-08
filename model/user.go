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
func CheckUser(name string) error {
	var user User
	err := db.Select("id").Where("username = ?", name).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ErrDBSelect
	}
	if user.ID > 0 {
		return errmsg.ErrDBNotUnique
	}
	return errmsg.Success
}

// CreateUser 添加用户
func CreateUser(u *User) errmsg.Code {
	err := db.Create(u).Error
	if err != nil {
		return errmsg.ErrDBInsert
	}
	return errmsg.Success
}

// IndexUser 查询用户列表
func IndexUser(perPage, page int, users []*User) errmsg.Code {
	err := db.Limit(perPage).Offset(OffsetByPage(perPage, page)).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ErrDBSelect
	}
	return errmsg.Success
}

// DeleteUser 删除用户
func DeleteUser(id int) errmsg.Code {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ErrDBDelete
	}
	return errmsg.Success
}

// UpdateUser 更新用户
func UpdateUser(id int, u *User) errmsg.Code {
	m := make(map[string]interface{})
	m["username"] = u.Username
	m["role"] = u.Role
	err := db.Model(&User{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errmsg.ErrDBUpdate
	}
	return errmsg.Success
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
	return ret
}
