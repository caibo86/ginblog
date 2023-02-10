package model

import (
	"github.com/caibo86/ginblog/api/base"
	"github.com/caibo86/ginblog/utils/crypt"
	"github.com/caibo86/ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);unique;not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(128);not null" json:"password" validate:"required,min=6,max=12" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2;not null" json:"role" validate:"required,gte=2" label:"角色码"`
}

// CheckUserExist 查询用户是否已经存在
func CheckUserExist(name string) (bool, error) {
	user := &User{}
	err := db.Select("id").Where("username = ?", name).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// CreateUser 添加用户
func CreateUser(u *User) error {
	return db.Create(u).Error
}

// IndexUser 查询用户列表
func IndexUser(perPage, page int) ([]*User, int64, error) {
	var users []*User
	var total int64
	err := db.Limit(perPage).Offset(base.OffsetByPage(perPage, page)).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return users, total, nil
}

// DeleteUser 删除用户
func DeleteUser(id int) error {
	return db.Where("id = ?", id).Delete(&User{}).Error
}

// UpdateUser 更新用户
func UpdateUser(id int, u *User) error {
	u.ID = uint(id)
	return db.Model(u).Select("username", "role").Updates(u).Error
}

// CheckLogin 登录验证
func CheckLogin(username, password string) error {
	user := &User{}
	err := db.Where("username = ?", username).First(user).Error
	if err != nil {
		return err
	}
	if crypt.Scrypt(password) != user.Password {
		return errmsg.ErrPasswordWrong
	}
	if user.Role != 1 {
		return errmsg.ErrNoPermission
	}
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	u.Password = crypt.Scrypt(u.Password)
	return nil
}
