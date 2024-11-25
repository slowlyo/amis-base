package models

import base "amis-base/internal/models"

type AdminPage struct {
	base.BaseModel

	Name   string `gorm:"type:varchar(255);not null;unique"`
	Sign   string `gorm:"type:varchar(255);not null;unique"`
	Schema string `gorm:"type:longtext"`
}
