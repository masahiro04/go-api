package dao

import "gorm.io/gorm"

type RW struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *RW {
	return &RW{
		DB: db,
	}
}
