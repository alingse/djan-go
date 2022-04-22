package djan

import "gorm.io/gorm"

type djanAdmin struct {
	db *gorm.DB
}

func NewAdmin(db *gorm.DB) *djanAdmin {
	return &djanAdmin{
		db: db,
	}
}
