package model

import "gorm.io/gorm"

// Profile 用户信息
type Profile struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20)" json:"name"`
	Desc   string `gorm:"type:varchar(200)" json:"desc"`
	QQChat string `gorm:"type:varchar(200)" json:"qq_chat"`
	WeChat string `gorm:"type:varchar(200)" json:"wechat"`
	Weibo  string `gorm:"type:varchar(200)" json:"weibo"`
	Bili   string `gorm:"type:varchar(200)" json:"bili"`
	Email  string `gorm:"type:varchar(200)" json:"email"`
	Image  string `gorm:"type:varchar(200)" json:"image"`
	Avatar string `gorm:"type:varchar(200)" json:"avatar"`
}

// ShowProfile 获取个人信息
func ShowProfile(id int) (*Profile, error) {
	profile := &Profile{}
	err := db.Where("id = ?", id).First(profile).Error
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// UpdateProfile 更改个人信息
func UpdateProfile(id int, profile *Profile) error {
	return db.Model(profile).Where("id= ?", id).Updates(profile).Error
}
