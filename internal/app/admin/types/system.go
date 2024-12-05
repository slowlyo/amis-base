package types

// SuperAdminSign 超级管理员标识
var SuperAdminSign = "administrator"

// AdminRouteMeta 后台路由元信息
type AdminRouteMeta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Hide  bool   `json:"hide"`
	Sort  int    `json:"sort"`
}

// AdminRoutes 后台路由
type AdminRoutes struct {
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	PageSign  string         `json:"pageSign"`
	IsFull    int            `json:"isFull"`
	IsHome    int            `json:"isHome"`
	Component string         `json:"component"`
	Meta      AdminRouteMeta `json:"meta"`
	Children  *[]AdminRoutes `json:"children"`
}

// Options 下拉框
type Options struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
