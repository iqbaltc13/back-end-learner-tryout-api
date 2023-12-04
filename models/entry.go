package models

import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  string
}

func (entry *Entry) Save() (*Entry, error) {

	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
