import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"time"

	"gorm.io/gorm"
)

type MateriCategory struct {
	gorm.Model
	ID           string         `gorm:"primaryKey"`
	Name         string         `gorm:"size:1000;not null;column:name" json:"name"`
	CreatedAt    string         `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt    string         `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt    string         `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
	MasterMateri []MasterMateri `gorm:"foreignKey:Category;references:ID"`
}

type MasterMateri struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	TypeID   string `gorm:"size:1000;not null;column:type_id" json:"type_id"`
	ClassID  string `gorm:"size:1000;not null;column:class_id" json:"class_id"`
	Category string `gorm:"not null;column:category" json:"category"`

	LinkFile   string     `gorm:"not null;column:link_file" json:"link_file"`
	Name       string     `gorm:"not null;column:name" json:"name"`
	Status     integer    `gorm:"not null;column:status" json:"status"`
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