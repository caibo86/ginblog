package model

import (
	"github.com/caibo86/ginblog/api/base"
	"gorm.io/gorm"
)

type Article struct {
	Category Category
	gorm.Model
	Title      string `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID int    `gorm:"type:int;not null" json:"category_id"`
	Desc       string `gorm:"type:varchar(200)" json:"desc"`
	Content    string `gorm:"type:longtext" json:"content"`
	Image      string `gorm:"type:varchar(100)" json:"image"`
}

// CreateArticle 创建文章
func CreateArticle(a *Article) error {
	return db.Create(a).Error
}

// IndexArticle 查询文章
func IndexArticle(perPage, page int, categoryID int) ([]*Article, int64, error) {
	var articles []*Article
	var total int64
	var err error
	if categoryID > 0 {
		err = db.Preload("Category").Limit(perPage).Offset(base.OffsetByPage(perPage, page)).Where("category_id = ?", categoryID).Find(&articles).Count(&total).Error
	} else {
		err = db.Preload("Category").Limit(perPage).Offset(base.OffsetByPage(perPage, page)).Find(&articles).Count(&total).Error
	}
	if err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}

// ShowArticle 查询单个文章
func ShowArticle(id int) (*Article, error) {
	a := &Article{}
	err := db.Preload("Category").Where("id = ?", id).First(a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

// UpdateArticle 更新文章
func UpdateArticle(id int, a *Article) error {
	a.ID = uint(id)
	return db.Model(a).Select("title", "category_id", "desc", "content", "image").Updates(a).Error
}

// DeleteArticle 删除文章
func DeleteArticle(id int) error {
	return db.Delete(&Article{}, "id = ?", id).Error
}
