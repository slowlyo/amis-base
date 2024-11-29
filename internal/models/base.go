package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime;not null;<-:create" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;not null" json:"updated_at"`
}

type BaseModelWithDeletedAt struct {
	BaseModel
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
