package models

import (
	base "amis-base/internal/models"
	"encoding/json"
)

type AdminPage struct {
	base.BaseModel

	Name   string          `gorm:"type:varchar(255);not null;unique" json:"name"`
	Sign   string          `gorm:"type:varchar(255);not null;unique" json:"sign"`
	Schema json.RawMessage `gorm:"type:longtext" json:"schema"`
}
