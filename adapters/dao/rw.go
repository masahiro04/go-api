package dao

import "gorm.io/gorm"

type RW struct {
	db *gorm.DB
}

func New(db *gorm.DB) *RW {
	return &RW{
		db: db,
	}
}
