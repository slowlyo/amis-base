package services

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type baseService struct{}

func (s baseService) BuildListPageQuery(query *gorm.DB, filters fiber.Map) *gorm.DB {
	// 分页
	query.Offset((filters["page"].(int) - 1) * filters["perPage"].(int)).Limit(filters["perPage"].(int))

	// 排序
	orderBy := filters["orderBy"].(string)
	orderDir := filters["orderDir"].(string)

	if orderBy != "" {
		query.Order(orderBy + " " + orderDir)
	}

	return query
}
