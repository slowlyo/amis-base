package models

type AdminPage struct {
	BaseModel

	Name   string `gorm:"type:varchar(255);not null;unique"`
	Sign   string `gorm:"type:varchar(255);not null;unique"`
	Schema string `gorm:"type:longtext"`
}
