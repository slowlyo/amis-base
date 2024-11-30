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

	db.Create(&[]models.AdminPage{
		{
			Name:   "页面管理",
			Sign:   "admin_page",
			Schema: json.RawMessage(`{ "type": "page", "className": "m:overflow-auto", "body": { "type": "crud", "perPage": 20, "filterTogglable": false, "filterDefaultVisible": false, "api": "/system/pages", "bulkActions": [ { "type": "button", "actionType": "dialog", "label": "删除", "icon": "fa-solid fa-trash-can", "dialog": { "type": "dialog", "title": "删除", "className": "py-2", "actions": [ { "type": "action", "actionType": "cancel", "label": "取消" }, { "type": "action", "actionType": "submit", "label": "删除", "level": "danger" } ], "body": [ { "type": "form", "wrapWithPanel": false, "api": { "method": "post", "url": "/system/pages", "data": { "id": "${ids}" } }, "body": [ { "type": "tpl", "className": "py-2", "tpl": "确认删除选中项？" } ] } ] } } ], "footerToolbar": [ "statistics", "pagination" ], "headerToolbar": [ { "type": "button", "actionType": "dialog", "dialog": { "type": "dialog", "title": "新增", "body": { "type": "form", "panelClassName": "px-48m:px-0", "title": "", "promptPageLeave": true, "onEvent": [], "body": [ { "type": "input-text", "name": "name", "label": "名称", "required": true }, { "type": "input-text", "name": "sign", "label": "标识", "required": true }, { "type": "input-sub-form", "name": "page", "label": "页面结构", "form": { "type": "form", "className": "h-full", "size": "full", "title": "", "body": { "type": "custom-amis-editor", "name": "schema", "label": "", "mode": "normal", "className": "h-full" } }, "required": true } ], "canAccessSuperData": false, "api": "post:/system/pages" }, "size": "md" }, "label": "新增", "icon": "fa fa-add", "level": "primary" }, "bulkActions", { "type": "reload", "align": "right" }, { "type": "filter-toggler", "align": "right" } ], "primaryField": "id", "columns": [ { "name": "id", "label": "ID", "sortable": true }, { "name": "name", "label": "名称", "searchable": true }, { "name": "sign", "label": "标识", "searchable": true }, { "name": "updated_at", "label": "更新时间", "type": "datetime", "sortable": true }, { "type": "operation", "label": "操作", "buttons": [ { "type": "button", "actionType": "dialog", "dialog": { "type": "dialog", "title": "编辑", "body": { "type": "form", "panelClassName": "px-48m:px-0", "title": "", "promptPageLeave": true, "onEvent": [], "body": [ { "type": "input-text", "name": "name", "label": "名称", "required": true }, { "type": "input-text", "name": "sign", "label": "标识", "required": true }, { "type": "input-sub-form", "name": "page", "label": "页面结构", "form": { "type": "form", "className": "h-full", "size": "full", "title": "", "body": { "type": "custom-amis-editor", "name": "schema", "label": "", "mode": "normal", "className": "h-full" } }, "required": true } ], "api": "put:/system/pages", "initApi": "/system/pages/edit?id=${id}", "redirect": "" }, "size": "md" }, "label": "编辑", "level": "link" }, { "type": "button", "actionType": "dialog", "label": "删除", "level": "link", "className": "text-danger", "dialog": { "type": "dialog", "title": "", "className": "py-2", "actions": [ { "type": "action", "actionType": "cancel", "label": "取消" }, { "type": "action", "actionType": "submit", "label": "删除", "level": "danger" } ], "body": [ { "type": "form", "wrapWithPanel": false, "api": "delete:/system/pages", "body": [ { "type": "tpl", "className": "py-2", "tpl": "确认删除选中项？" } ] } ] } } ] } ] } }`),
		},
		{
			Name:   "角色",
			Sign:   "admin_role",
			Schema: json.RawMessage(`{"type":"page","regions":["body"],"id":"u:a8979bbc7044","pullRefresh":{"disabled":true},"body":[{"type":"panel","title":"","body":[{"type":"tpl","tpl":"","wrapperComponent":"","inline":false,"id":"u:921ea986c7c6"},{"id":"u:ebed804f316f","type":"crud2","mode":"table2","dsType":"api","syncLocation":true,"selectable":true,"multiple":true,"primaryField":"id","loadType":"pagination","api":{"url":"/system/roles","method":"get"},"quickSaveItemApi":{"url":"/system/roles","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"filter":{"type":"form","title":"筛选","mode":"inline","columnCount":4,"clearValueOnHidden":true,"behavior":["SimpleQuery"],"body":[{"name":"name","label":"名称","type":"input-text","size":"full","required":false,"behavior":"SimpleQuery","id":"u:e6ad3142299a","clearable":true},{"name":"sign","label":"标识","type":"input-text","size":"full","required":false,"behavior":"SimpleQuery","id":"u:0f93933f2010","clearable":true}],"actions":[{"type":"reset","label":"重置","id":"u:e5e0860c35e7"},{"type":"submit","label":"查询","level":"primary","id":"u:dfddb6d9b845"}],"id":"u:625a8a01302f","feat":"Insert","autoFocus":false,"wrapWithPanel":true,"affixFooter":true,"themeCss":{"headerControlClassName":{}},"visibleOn":"${showFilter}","labelAlign":"left"},"headerToolbar":[{"type":"flex","direction":"row","justify":"flex-start","alignItems":"stretch","style":{"position":"static"},"items":[{"type":"container","align":"left","behavior":["Insert","BulkEdit","BulkDelete"],"body":[{"type":"button","label":"新增","level":"primary","className":"m-r-xs","behavior":"Insert","onEvent":{"click":{"actions":[{"actionType":"dialog","dialog":{"type":"dialog","body":[{"id":"u:077351270b41","type":"form","title":"新增数据","mode":"flex","labelAlign":"top","dsType":"api","feat":"Insert","body":[{"name":"name","label":"名称","row":0,"type":"input-text","id":"u:9f939987f6cb","clearable":true,"required":true,"validations":{"maxLength":255}},{"name":"sign","label":"标识","row":1,"type":"input-text","id":"u:2ef00c8e400d","required":true,"validations":{"maxLength":255},"clearable":true}],"api":{"url":"/system/roles","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"resetAfterSubmit":true,"actions":[{"type":"button","actionType":"cancel","label":"取消"},{"type":"button","actionType":"submit","label":"提交","level":"primary"}],"onEvent":{"submitSucc":{"actions":[{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f"}]}}}],"title":"新增数据","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:753bbec4a163"},{"type":"button","actionType":"submit","label":"提交","level":"primary","id":"u:a6781c3c01bd"}],"actionType":"dialog","id":"u:fe8bcc856a05","showCloseButton":true,"closeOnOutside":false,"closeOnEsc":false,"showErrorMsg":true,"showLoading":true,"draggable":true}}]}},"id":"u:be3d0c329604","icon":"fa fa-plus"},{"type":"button","label":"删除","behavior":"BulkDelete","level":"danger","className":"m-r-xs","confirmText":"确认要批量删除数据「${JOIN(ARRAYMAP(selectedItems, item => item.id), ',')}」","disabledOn":"${selectedItems != null && selectedItems.length < 1}","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"method":"post","url":"/system/roles/delete","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${ids}"}}},{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f","args":{},"expression":"${responseResult.responseStatus == 0}"}]}},"id":"u:f6976a761eab","icon":"fa fa-trash-o"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexBasis":"auto","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-start"},"id":"u:d2feb4775db8"},{"type":"container","align":"right","behavior":["FuzzyQuery"],"body":[{"type":"button","label":"筛选","onEvent":{"click":{"actions":[{"componentId":"u:a8979bbc7044","ignoreError":false,"actionType":"setValue","args":{"value":{"showFilter":"${!showFilter}"}}}]}},"id":"u:341a73d9cbf1","icon":"fa fa-search","className":{"text-primary":"${showFilter}","border-primary":"${showFilter}"}},{"type":"button","label":"","onEvent":{"click":{"actions":[{"componentId":"u:ebed804f316f","ignoreError":false,"outputVar":"","actionType":"reload"}]}},"id":"u:405b59ca63aa","icon":"fa fa-refresh","className":"m-l-xs"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-end"},"id":"u:9de66d99d8ea","isFixedHeight":false}],"id":"u:b2da8147240d"}],"footerToolbar":[{"type":"flex","direction":"row","justify":"flex-start","alignItems":"stretch","style":{"position":"static"},"items":[{"type":"container","align":"left","body":[],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-start"},"id":"u:4827bfdfabb9","isFixedHeight":false},{"type":"container","align":"right","body":[{"type":"pagination","behavior":"Pagination","layout":["total","perPage","pager","go"],"perPage":20,"perPageAvailable":[20,50,100,200],"align":"right","id":"u:219881e04744","size":"md"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexBasis":"auto","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-end"},"id":"u:6534f7aa4d15"}],"id":"u:39349b5d361d"}],"columns":[{"type":"tpl","title":"ID","name":"id","id":"u:f5854eff0cfa"},{"type":"tpl","title":"名称","name":"name","id":"u:315b256293b8"},{"type":"tpl","title":"标识","name":"sign","id":"u:204471470466","placeholder":"-"},{"type":"datetime","format":"YYYY-MM-DD HH:mm:ss","value":1732889273,"name":"updated_at","title":"更新时间","id":"u:004dde22d187","placeholder":"-","sorter":true},{"type":"operation","title":"操作","buttons":[{"type":"button","label":"编辑","level":"link","behavior":"Edit","onEvent":{"click":{"actions":[{"actionType":"dialog","dialog":{"type":"dialog","body":[{"id":"u:213c44c60d36","type":"form","title":"编辑数据","mode":"flex","labelAlign":"top","dsType":"api","feat":"Edit","body":[{"name":"id","label":"id","row":0,"type":"input-number","id":"u:5736f77c8961","keyboard":true,"step":1,"visible":false,"hidden":true},{"name":"name","label":"名称","row":1,"type":"input-text","id":"u:841f35203ece","clearable":true,"required":true,"validations":{"maxLength":255}},{"name":"sign","label":"标识","row":2,"type":"input-text","id":"u:030699b39cac","required":true,"validations":{"maxLength":255},"clearable":true}],"api":{"url":"/system/roles","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"resetAfterSubmit":true,"initApi":{"url":"/system/roles/detail","method":"get","requestAdaptor":"","adaptor":"","messages":{},"data":{"id":"${id}"}},"actions":[{"type":"button","actionType":"cancel","label":"取消"},{"type":"button","actionType":"submit","label":"提交","level":"primary"}],"onEvent":{"submitSucc":{"actions":[{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f"}]}}}],"title":"编辑数据","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:97cbeb046707"},{"type":"button","actionType":"submit","label":"提交","level":"primary","id":"u:60ae7a87557d"}],"actionType":"dialog","id":"u:61abf5f368fa","showCloseButton":true,"closeOnOutside":false,"closeOnEsc":false,"showErrorMsg":true,"showLoading":true,"draggable":true}}]}},"id":"u:7cae73c66cd0"},{"type":"button","label":"删除","behavior":"Delete","className":"m-r-xs text-danger","level":"link","confirmText":"确认要删除数据","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"method":"post","url":"/system/roles/delete","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${TRIM(id)}"}},"outputVar":"responseResult","options":{}},{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f","args":{"query":""},"expression":"${responseResult.responseStatus == 0}"}]}},"id":"u:df00563b639d"}],"id":"u:8136b229cca9"}],"editorSetting":{"mock":{"enable":true,"maxDisplayRows":5}}}],"id":"u:93deb9fe8588","affixFooter":false,"actions":[]}],"definitions":{},"data":{"showFilter":false}}`),
		},
	})
}
