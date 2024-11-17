package model

import (
	"time"

	"gorm.io/gorm"
)

func MakeMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
	)
}

type Model struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
