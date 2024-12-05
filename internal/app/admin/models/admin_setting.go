package models

import base "amis-base/internal/models"

type AdminSetting struct {
	base.BaseModel

	Key   string `gorm:"type:varchar(255);not null;unique_index" json:"key"`
	Value string `gorm:"type:mediumtext" json:"value"`
}
