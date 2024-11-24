package models

type AdminPermission struct {
	BaseModel

	ParentId uint   `gorm:"type:int(11);not null;default:0"`
	Name     string `gorm:"type:varchar(255);not null;unique"`
	Sign     string `gorm:"type:varchar(255);not null;unique"`
	Api      string `gorm:"type:text"`
	Sort     uint   `gorm:"type:int(11);not null;default:0"`

	Roles []AdminRole `gorm:"many2many:admin_role_permission;"`
	Menus []AdminMenu `gorm:"many2many:admin_menu_permission;"`
}
