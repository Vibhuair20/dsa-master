package models

import (
	"gorm.io/gorm"
)

type topics struct {
	ID          uint `gorm:"primary key;autoIncrement" json:"id"`
	name        *string
	difficultye *string `gorm:"easy";difficultye`
	difficultym *string `gorm:"medium"; json:"difficultym"`
	difficultyh *string `gorm:"hard";json:"difficultyh"`
	selected    *string `gorm:"T/F"`
}

func MigrateTopics(db *gorm.DB) error {
	err := db.AutoMigrate(&topics{})
	return err
}
