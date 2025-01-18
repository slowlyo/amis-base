package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/cache"
	"amis-base/internal/pkg/db"
	"amis-base/internal/pkg/helper"
)

type AdminSetting struct {
}

func (s AdminSetting) Set(key string, value any) error {
	var record models.AdminSetting

	result := db.Query().Where("`key` = ?", key).Find(&record)

	record.Key = key
	record.Value = helper.JsonEncode(value)

	_ = cache.Delete(cache.BuildKey("setting:" + key))

	if result.RowsAffected == 0 {
		return db.Query().Create(&record).Error
	}

	return db.Query().Save(&record).Error
}

func (s AdminSetting) Get(key string) any {
	value, err := cache.RememberForever[any](cache.BuildKey("setting:"+key), func() (any, error) {
		var record models.AdminSetting

		result := db.Query().Where("`key` = ?", key).First(&record)

		if result.RowsAffected == 0 {
			return "", nil
		}

		return helper.JsonDecode[any](record.Value), nil
	})

	if err != nil {
		return ""
	}

	return value
}
