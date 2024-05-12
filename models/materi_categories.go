package models

import (
	"fmt"

	"github.com/iqbaltc13/back-end-learner-tryout-api/database"
	"gorm.io/gorm"
)

type MateriCategory struct {
	gorm.Model
	ID             string         `gorm:"primaryKey"`
	Name           string         `gorm:"size:1000;not null;column:name" json:"name"`
	CreatedAt      string         `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt      string         `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt      string         `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
	MasterMateries []MasterMateri `gorm:"foreignKey:Category;references:ID"`
}

type MasterMateri struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	TypeID   string `gorm:"size:1000;not null;column:type_id" json:"type_id"`
	ClassID  string `gorm:"size:1000;not null;column:class_id" json:"class_id"`
	Category string `gorm:"not null;column:category" json:"category"`

	LinkFile   string     `gorm:"not null;column:link_file" json:"link_file"`
	Name       string     `gorm:"not null;column:name" json:"name"`
	Status     int        `gorm:"not null;column:status" json:"status"`
	CreatedAt  string     `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt  string     `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt  string     `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
	MasterType MateriType `gorm:"foreignKey:ID;references:TypeID"`
}

type MateriType struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"size:1000;not null;column:name" json:"name"`

	CreatedAt string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (MasterMateri) TableName() string {
	return "master_materies"
}

func FindMasterMateriesPerCategoriesByClassIds(ids []string) ([]MateriCategory, error) {
	var materiCategories []MateriCategory
	fmt.Println(ids)
	err := database.Database.Preload("MasterMateries", func(db *gorm.DB) *gorm.DB {
		return db.Where("class_id IN ?", ids).Preload("MateriType").Group("master_materies.type_id")
	}).Find(&materiCategories).Error
	if err != nil {
		return materiCategories, err
	}
	return materiCategories, nil
}
