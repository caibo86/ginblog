package model

import (
	"github.com/caibo86/ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(128);not null" json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) error {
	c := &Category{}
	err := db.Select("id").Where("name = ?", name).First(c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ErrDBSelect
	}
	if c.ID > 0 {
		return errmsg.ErrDBNotUnique
	}
	return errmsg.Success
}

// IndexCategory 查询分类
func IndexCategory(perPage, page int) ([]*Category, error) {
	var categories []*Category
	err := db.Limit(perPage).Offset(OffsetByPage(perPage, page)).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ErrDBSelect
	}
	return categories, errmsg.Success
}

// CreateCategory 创建分类
func CreateCategory(c *Category) error {
	code := CheckCategory(c.Name)
	if code != errmsg.Success {
		return code
	}
	err := db.Create(c).Error
	if err != nil {
		return errmsg.ErrDBInsert
	}
	return errmsg.Success
}

// DeleteCategory 删除分类
func DeleteCategory(id int) error {
	err := db.Delete(&Category{}, "id = ?", id).Error
	if err != nil {
		return errmsg.ErrDBDelete
	}
	return errmsg.Success
}

// UpdateCategory 更新分类
func UpdateCategory(id int, c *Category) error {
	code := CheckCategory(c.Name)
	if code != errmsg.Success {
		return code
	}
	c.ID = id
	err := db.Model(c).Select("name").Updates(c).Error
	if err != nil {
		return errmsg.ErrDBUpdate
	}
	return errmsg.Success
}

// OffsetByPage 根据页计算偏移
func OffsetByPage(perPage, page int) int {
	return (page - 1) * perPage
}
