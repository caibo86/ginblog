package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	CID     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Image   string `gorm:"type:varchar(100)" json:"image"`
}
