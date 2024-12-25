package services

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type baseService struct{}

func (s baseService) ListPaginate(query *gorm.DB, filters fiber.Map) *gorm.DB {
	// 分页
	query = query.Offset((filters["page"].(int) - 1) * filters["perPage"].(int)).Limit(filters["perPage"].(int))

	// 排序
	query = s.Sortable(query, filters)

	return query
}

func (s baseService) ListGet(query *gorm.DB, filters fiber.Map) *gorm.DB {
	// 排序
	query = s.Sortable(query, filters)

	return query
}

func (s baseService) Sortable(query *gorm.DB, filters fiber.Map) *gorm.DB {
	orderBy := filters["orderBy"].(string)
	orderDir := filters["orderDir"].(string)

	if orderBy != "" {
		query.Order(orderBy + " " + orderDir)
	}

	return query
}
