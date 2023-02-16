package model

import (
	"github.com/caibo86/ginblog/api/base"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(128);unique;not null" json:"name"`
}

// CheckCategoryExist 查询分类是否已经存在
func CheckCategoryExist(name string) (bool, error) {
	c := &Category{}
	err := db.Select("id").Where("name = ?", name).First(c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if c.ID > 0 {
		return true, nil
	}
	return false, nil
}

// IndexCategory 查询分类
func IndexCategory(pageSize, page int) ([]*Category, int64, error) {
	var categories []*Category
	var total int64
	err := db.Model(&Category{}).Count(&total).Limit(pageSize).Offset(base.OffsetByPage(pageSize, page)).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return categories, total, nil
}

// CreateCategory 创建分类
func CreateCategory(c *Category) error {
	return db.Create(c).Error
}

// DeleteCategory 删除分类
func DeleteCategory(id int) error {
	return db.Delete(&Category{}, "id = ?", id).Error
}

// UpdateCategory 更新分类
func UpdateCategory(id int, c *Category) error {
	c.ID = id
	return db.Model(c).Select("name").Updates(c).Error
}
