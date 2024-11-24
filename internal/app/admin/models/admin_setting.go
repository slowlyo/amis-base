package models

type AdminSetting struct {
	BaseModel

	Key   string `gorm:"type:varchar(255);not null;unique_index"`
	Value string `gorm:"type:mediumtext"`
}
