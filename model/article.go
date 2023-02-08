package model

import (
	"github.com/caibo86/ginblog/utils/errmsg"
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
	err := db.Create(a).Error
	if err != nil {
		return errmsg.ErrDBInsert
	}
	return errmsg.Success
}

// IndexArticle 查询文章
func IndexArticle(perPage, page int) ([]*Article, error) {
	var articles []*Article
	err := db.Limit(perPage).Offset(OffsetByPage(perPage, page)).Find(&articles).Error
	if err != nil {
		return nil, errmsg.ErrDBSelect
	}
	return articles, errmsg.Success
}

// ShowArticle 查询单个文章
func ShowArticle(id int) (*Article, error) {
	var a *Article
	return a, errmsg.Success
}

// UpdateArticle 更新文章
func UpdateArticle(id int, a *Article) error {
	a.ID = uint(id)
	if err := db.Model(a).Select("title", "cid", "desc", "content", "img").Updates(a).Error; err != nil {
		return errmsg.ErrDBUpdate
	}
	return errmsg.Success
}

// DeleteArticle 删除文章
func DeleteArticle(id int) error {
	if err := db.Delete(&Article{}, "id = ?", id).Error; err != nil {
		return errmsg.ErrDBDelete
	}
	return errmsg.Success
}
