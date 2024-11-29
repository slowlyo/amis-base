package services

type BaseServiceInterface interface {
	List(page, perPage int) ([]any, int64)
}

// service 的默认实现
type baseService struct{}
