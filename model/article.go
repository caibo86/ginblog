package model

import (
	"github.com/caibo86/ginblog/api/base"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title      string   `gorm:"type:varchar(100);not null" json:"title"`
	Desc       string   `gorm:"type:varchar(200)" json:"desc"`
	Content    string   `gorm:"type:longtext" json:"content"`
	Image      string   `gorm:"type:varchar(100)" json:"image"`
	CategoryID int      `gorm:"type:int;not null" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID"`
}

// CreateArticle 创建文章
func CreateArticle(a *Article) error {
	err := db.Create(a).Error
	if err == nil {
		db.Find(&a.Category, "id = ?", a.CategoryID)
	}
	return err
}

// IndexArticle 查询文章
func IndexArticle(pageSize, page int, categoryID int, title string) ([]*Article, int64, error) {
	var articles []*Article
	var total int64
	var err error

	if categoryID > 0 {
		if title != "" {
			err = db.Model(&Article{}).Where("category_id = ? and title LIKE ?", categoryID, title+"%").Count(&total).
				Preload("Category").Limit(pageSize).Offset(base.OffsetByPage(pageSize, page)).Find(&articles).Error
		} else {
			err = db.Model(&Article{}).Where("category_id = ?", categoryID).Count(&total).
				Preload("Category").Limit(pageSize).Offset(base.OffsetByPage(pageSize, page)).Find(&articles).Error
		}
	} else {
		if title != "" {
			err = db.Model(&Article{}).Where("title LIKE ?", title+"%").Count(&total).
				Preload("Category").Limit(pageSize).Offset(base.OffsetByPage(pageSize, page)).Find(&articles).Error
		} else {
			err = db.Model(&Article{}).Count(&total).
				Preload("Category").Limit(pageSize).Offset(base.OffsetByPage(pageSize, page)).Find(&articles).Error
		}

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
	err := db.Model(a).Select("title", "category_id", "desc", "content", "image").Updates(a).Error
	if err == nil {
		db.Find(&a.Category, "id = ?", a.CategoryID)
	}
	return err
}

// DeleteArticle 删除文章
func DeleteArticle(id int) error {
	return db.Delete(&Article{}, "id = ?", id).Error
}
