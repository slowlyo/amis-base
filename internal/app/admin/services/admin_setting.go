package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/db"
	"encoding/json"
)

type AdminSetting struct {
}

func (s AdminSetting) Set(key string, value any) error {
	var record models.AdminSetting

	result := db.GetDB().Where("`key` = ?", key).Find(&record)

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	record.Key = key
	record.Value = string(jsonValue)

	if result.RowsAffected == 0 {
		return db.GetDB().Create(&record).Error
	}

	return db.GetDB().Save(&record).Error
}

func (s AdminSetting) Get(key string) any {
	var record models.AdminSetting

	result := db.GetDB().Where("`key` = ?", key).First(&record)

	if result.RowsAffected == 0 {
		return ""
	}

	var value any

	err := json.Unmarshal([]byte(record.Value), &value)
	if err != nil {
		return ""
	}

	return value
}
