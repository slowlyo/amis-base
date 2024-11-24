package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type BaseModelWithDeletedAt struct {
	BaseModel
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
