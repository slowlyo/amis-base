package models

import "time"

type Token struct {
	BaseModel

	TableName  string    `gorm:"type:varchar(255);not null"`
	UserId     uint      `gorm:"type:int(11);not null;default:0;index"`
	Token      string    `gorm:"type:varchar(64);not null"`
	LastUsedAt time.Time `gorm:"type:datetime;not null"`
}
