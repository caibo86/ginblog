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
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(128);not null" json:"password" validate:"required,min=6,max=12" label:"密码"`
	Role     int    `gorm:"type:int;not null;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
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
	return errmsg.OK
}

// CreateUser 添加用户
func CreateUser(u *User) error {
	err := db.Create(u).Error
	if err != nil {
		return errmsg.ErrDBInsert
	}
	return errmsg.OK
}

// IndexUser 查询用户列表
func IndexUser(perPage, page int) ([]*User, int64, error) {
	var users []*User
	var total int64
	err := db.Limit(perPage).Offset(OffsetByPage(perPage, page)).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errmsg.ErrDBSelect
	}
	return users, total, errmsg.OK
}

// DeleteUser 删除用户
func DeleteUser(id int) error {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ErrDBDelete
	}
	return errmsg.OK
}

// UpdateUser 更新用户
func UpdateUser(id int, u *User) error {
	u.ID = uint(id)
	err := db.Model(u).Select("username", "role").Updates(u).Error
	if err != nil {
		return errmsg.ErrDBUpdate
	}
	return errmsg.OK
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

// CheckLogin 登录验证
func CheckLogin(username, password string) error {
	user := &User{}
	err := db.Where("username = ?", username).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return errmsg.ErrRecordNotFound
	} else if err != nil {
		return errmsg.ErrDBSelect
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ErrPasswordWrong
	}
	if user.Role != 1 {
		return errmsg.ErrNoPermission
	}
	return errmsg.OK
}
