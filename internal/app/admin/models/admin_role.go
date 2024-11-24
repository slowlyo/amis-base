package models

type AdminRole struct {
	BaseModel

	Name string `gorm:"type:varchar(255);not null;unique"`
	Sign string `gorm:"type:varchar(255);not null;unique"`

	Users       []AdminUser       `gorm:"many2many:admin_user_role;"`
	Permissions []AdminPermission `gorm:"many2many:admin_role_permission;"`
}
