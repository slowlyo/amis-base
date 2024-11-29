package db

import (
	"amis-base/internal/app/admin/models"
	baseModel "amis-base/internal/models"
	"amis-base/internal/pkg/helper"
	"encoding/json"
)

// Seed 填充初始数据
func Seed() {
	// 用户
	go seedUsers()

	// 菜单
	go seedMenus()

	// 角色
	go seedRoles()

	// 设置
	go seedSettings()

	// 页面
	go seedPages()
}

// 判断数据表是否为空
func isNull(model interface{}) bool {
	var count int64
	db.Model(&model).Count(&count)

	return count == 0
}

// 填充用户
func seedUsers() {
	if !isNull(models.AdminUser{}) {
		return
	}

	db.Create(&models.AdminUser{
		Name:     "Administrator",
		Username: "admin",
		Password: helper.BcryptString("admin"),
		Enabled:  1,
		Avatar:   "https://api.dicebear.com/9.x/bottts-neutral/png?seed=amis-base",
	})

	// 默认角色
	db.Table("admin_user_role").Create(&map[string]interface{}{
		"admin_user_id": 1,
		"admin_role_id": 1,
	})
}

// 填充菜单
func seedMenus() {
	if !isNull(models.AdminMenu{}) {
		return
	}

	db.Create(&[]models.AdminMenu{
		{
			BaseModel: baseModel.BaseModel{ID: 1},
			Name:      "控制台",
			Icon:      "ph:chart-line-up-fill",
			Path:      "/dashboard",
			IsHome:    1,
			Visible:   1,
			PageSign:  "dashboard",
		},
		{
			BaseModel: baseModel.BaseModel{ID: 2},
			Name:      "系统管理",
			Icon:      "material-symbols:settings-outline",
			Path:      "/system",
			Visible:   1,
		},
		{
			BaseModel: baseModel.BaseModel{ID: 3},
			ParentId:  2,
			Name:      "管理员",
			Icon:      "ph:user-gear",
			Path:      "/system/admin_user",
			Visible:   1,
			PageSign:  "admin_user",
		},
		{
			BaseModel: baseModel.BaseModel{ID: 4},
			ParentId:  2,
			Name:      "角色",
			Icon:      "carbon:user-role",
			Path:      "/system/admin_role",
			Visible:   1,
			PageSign:  "admin_role",
		},
		{
			BaseModel: baseModel.BaseModel{ID: 5},
			ParentId:  2,
			Name:      "权限",
			Icon:      "fluent-mdl2:permissions",
			Path:      "/system/admin_permission",
			Visible:   1,
			PageSign:  "admin_permission",
		},
		{
			BaseModel: baseModel.BaseModel{ID: 6},
			ParentId:  2,
			Name:      "菜单",
			Icon:      "ant-design:menu-unfold-outlined",
			Path:      "/system/admin_menu",
			Visible:   1,
			PageSign:  "admin_menu",
		},
		{
			BaseModel: baseModel.BaseModel{ID: 7},
			Name:      "个人中心",
			Icon:      "ph:circle",
			Path:      "/user",
			Visible:   0,
			PageSign:  "user",
		},
	})
}

// 填充角色
func seedRoles() {
	if !isNull(models.AdminRole{}) {
		return
	}

	db.Create(&[]models.AdminRole{
		{
			Name: "超级管理员",
			Sign: "administrator",
		},
	})
}

// 填充设置
func seedSettings() {
	if !isNull(models.AdminSetting{}) {
		return
	}

	db.Create(&models.AdminSetting{
		Key:   "system.theme",
		Value: `{"darkTheme":false,"footer":false,"breadcrumb":true,"themeColor":"rgb(22,119,255)","layoutMode":"default","siderTheme":"light","topTheme":"light","animateInType":"alpha","animateInDuration":600,"animateOutType":"alpha","animateOutDuration":600,"loginTemplate":"default","keepAlive":false,"enableTab":false,"tabIcon":false,"accordionMenu":false}`,
	})
}

// 填充页面
func seedPages() {
	if !isNull(models.AdminPage{}) {
		return
	}

	db.Create(&models.AdminPage{
		Name:   "页面管理",
		Sign:   "admin_page",
		Schema: json.RawMessage(`{ "type": "page", "className": "m:overflow-auto", "body": { "type": "crud", "perPage": 20, "filterTogglable": false, "filterDefaultVisible": false, "api": "/system/pages", "bulkActions": [ { "type": "button", "actionType": "dialog", "label": "删除", "icon": "fa-solid fa-trash-can", "dialog": { "type": "dialog", "title": "删除", "className": "py-2", "actions": [ { "type": "action", "actionType": "cancel", "label": "取消" }, { "type": "action", "actionType": "submit", "label": "删除", "level": "danger" } ], "body": [ { "type": "form", "wrapWithPanel": false, "api": { "method": "post", "url": "/system/pages", "data": { "id": "${ids}" } }, "body": [ { "type": "tpl", "className": "py-2", "tpl": "确认删除选中项？" } ] } ] } } ], "footerToolbar": [ "statistics", "pagination" ], "headerToolbar": [ { "type": "button", "actionType": "dialog", "dialog": { "type": "dialog", "title": "新增", "body": { "type": "form", "panelClassName": "px-48m:px-0", "title": "", "promptPageLeave": true, "onEvent": [], "body": [ { "type": "input-text", "name": "name", "label": "名称", "required": true }, { "type": "input-text", "name": "sign", "label": "标识", "required": true }, { "type": "input-sub-form", "name": "page", "label": "页面结构", "form": { "type": "form", "className": "h-full", "size": "full", "title": "", "body": { "type": "custom-amis-editor", "name": "schema", "label": "", "mode": "normal", "className": "h-full" } }, "required": true } ], "canAccessSuperData": false, "api": "post:/system/pages" }, "size": "md" }, "label": "新增", "icon": "fa fa-add", "level": "primary" }, "bulkActions", { "type": "reload", "align": "right" }, { "type": "filter-toggler", "align": "right" } ], "primaryField": "id", "columns": [ { "name": "id", "label": "ID", "sortable": true }, { "name": "name", "label": "名称", "searchable": true }, { "name": "sign", "label": "标识", "searchable": true }, { "name": "updated_at", "label": "更新时间", "type": "datetime", "sortable": true }, { "type": "operation", "label": "操作", "buttons": [ { "type": "button", "actionType": "dialog", "dialog": { "type": "dialog", "title": "编辑", "body": { "type": "form", "panelClassName": "px-48m:px-0", "title": "", "promptPageLeave": true, "onEvent": [], "body": [ { "type": "input-text", "name": "name", "label": "名称", "required": true }, { "type": "input-text", "name": "sign", "label": "标识", "required": true }, { "type": "input-sub-form", "name": "page", "label": "页面结构", "form": { "type": "form", "className": "h-full", "size": "full", "title": "", "body": { "type": "custom-amis-editor", "name": "schema", "label": "", "mode": "normal", "className": "h-full" } }, "required": true } ], "api": "put:/system/pages", "initApi": "/system/pages/edit?id=${id}", "redirect": "" }, "size": "md" }, "label": "编辑", "level": "link" }, { "type": "button", "actionType": "dialog", "label": "删除", "level": "link", "className": "text-danger", "dialog": { "type": "dialog", "title": "", "className": "py-2", "actions": [ { "type": "action", "actionType": "cancel", "label": "取消" }, { "type": "action", "actionType": "submit", "label": "删除", "level": "danger" } ], "body": [ { "type": "form", "wrapWithPanel": false, "api": "delete:/system/pages", "body": [ { "type": "tpl", "className": "py-2", "tpl": "确认删除选中项？" } ] } ] } } ] } ] } }`),
	})
}
