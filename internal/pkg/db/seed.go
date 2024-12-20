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

// ResetPages 重置页面
func ResetPages() {
	seedPages(true)
}

// 填充页面
func seedPages(force ...bool) {
	// 是否覆盖数据
	isForce := len(force) > 0 && force[0]

	if !isNull(models.AdminPage{}) && !isForce {
		return
	}

	if isForce {
		db.Model(&models.AdminPage{}).Exec("TRUNCATE TABLE admin_pages")
	}

	db.Create(&[]models.AdminPage{
		{
			Name:   "页面管理",
			Sign:   "admin_page",
			Schema: json.RawMessage(`{"body":[{"alwaysShowPagination":true,"api":{"method":"get","url":"/system/pages"},"bulkActions":[{"actionType":"ajax","api":"","confirmText":"确定要删除？","editorSetting":{"behavior":"bulkDelete"},"icon":"fa fa-trash-o","id":"u:394fe8eeb9b7","label":"删除","level":"danger","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${ids}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/pages/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button"}],"columns":[{"id":"u:154e022f601d","label":"ID","name":"id","placeholder":"-","sortable":true,"type":"text"},{"id":"u:7172c1ba66c8","label":"名称","name":"name","placeholder":"-","type":"text"},{"id":"u:acc2fc18bb5d","label":"标识","name":"sign","type":"text"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:05ee8b110a15","label":"创建时间","name":"created_at","placeholder":"-","sortable":true,"type":"date"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:a17c366510d7","label":"更新时间","name":"updated_at","placeholder":"-","sortable":true,"type":"date"},{"buttons":[{"actionType":"dialog","dialog":{"actionType":"dialog","actions":[{"actionType":"cancel","id":"u:114404b175c8","label":"取消","type":"button"},{"actionType":"submit","id":"u:b86dc1e4380b","label":"提交","level":"primary","type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/pages"},"body":[{"hidden":true,"id":"u:e5e6c2902816","keyboard":true,"label":"id","name":"id","row":0,"step":1,"type":"input-number","visible":false},{"clearable":true,"id":"u:3f288cda43d2","label":"名称","name":"name","required":true,"row":1,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:46707e48fcd8","label":"标识","name":"sign","required":true,"row":2,"type":"input-text","validations":{"maxLength":255}},{"form":{"body":{"className":"h-full","label":"","mode":"normal","name":"schema","type":"custom-amis-editor"},"className":"h-full","size":"full","title":"","type":"form"},"id":"u:624be988419b","label":"页面结构","name":"page","required":true,"type":"input-sub-form"}],"dsType":"api","feat":"Edit","id":"u:7c52686cc8e9","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/pages/detail"},"labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"}]}},"resetAfterSubmit":true,"title":"编辑数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"id":"u:1cc4b9e5e697","showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"dialog"},"editorSetting":{"behavior":"update"},"id":"u:29435e16d2df","label":"编辑","level":"link","type":"button"},{"behavior":"Edit","id":"u:7cae73c66cd0","label":"复制","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/pages/copy"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","args":{},"componentId":"u:563375ef114c","expression":"${responseResult.responseStatus == 0}","groupType":"component"}]}},"type":"button"},{"id":"u:97b94e051b85","label":"JSON","level":"link","onEvent":{"click":{"actions":[{"actionType":"dialog","dialog":{"$ref":"modal-ref-1"},"ignoreError":false}]}},"type":"button"},{"actionType":"","api":{"method":"post","url":"/system/pages"},"className":"text-danger","confirmText":"确定要删除？","editorSetting":{"behavior":"delete"},"id":"u:cee85dad465d","label":"删除","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${TRIM(id)}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/pages/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button"}],"id":"u:01259e435906","label":"操作","placeholder":"-","type":"operation"}],"filter":{"actions":[{"actionType":"clear-and-submit","id":"u:cdb68b8d4b45","label":"重置","type":"button"},{"id":"u:253f267c0d95","label":"搜索","level":"primary","type":"submit"}],"body":[{"clearable":true,"id":"u:7fed8045c4d1","label":"名称","name":"name","size":"md","type":"input-text"},{"clearable":true,"id":"u:7fed8045c4d1","label":"标识","name":"sign","size":"md","type":"input-text"}],"feat":"Insert","id":"u:ba2b6c5f28ea","panelClassName":"base-filter","title":"","type":"form"},"filterDefaultVisible":false,"filterTogglable":true,"footerToolbar":[{"type":"statistics"},{"type":"pagination"},{"body":[{"id":"u:8378ce2f122e","multiple":false,"name":"perPage","onEvent":{"change":{"actions":[{"actionType":"reload","componentId":"admin-api.system.admin_users.crud","data":{"perPage":"${event.data.value}"}}]}},"options":[{"label":"20 条/页","value":20},{"label":"30 条/页","value":30},{"label":"50 条/页","value":50},{"label":"100 条/页","value":100},{"label":"200 条/页","value":200}],"overlayPlacement":"top","selectFirst":true,"type":"select"}],"feat":"Insert","id":"u:abf9235be81a","target":"window","type":"form","wrapWithPanel":false}],"headerToolbar":[{"actionType":"dialog","dialog":{"actionType":"dialog","actions":[{"actionType":"cancel","id":"u:4b83a945c64a","label":"取消","type":"button"},{"actionType":"submit","id":"u:947858a0fae4","label":"提交","level":"primary","type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/pages"},"body":[{"hidden":true,"id":"u:1248c8850238","keyboard":true,"label":"id","name":"id","row":0,"step":1,"type":"input-number","visible":false},{"clearable":true,"id":"u:7541fa023716","label":"名称","name":"name","required":true,"row":1,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:00fba10e51a0","label":"标识","name":"sign","required":true,"row":2,"type":"input-text","validations":{"maxLength":255}},{"form":{"body":{"className":"h-full","label":"","mode":"normal","name":"schema","type":"custom-amis-editor"},"className":"h-full","size":"full","title":"","type":"form"},"id":"u:954dc473afaf","label":"页面结构","name":"page","required":true,"type":"input-sub-form"}],"dsType":"api","feat":"Edit","id":"u:0269868efcdf","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/pages/detail"},"labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"}]}},"resetAfterSubmit":true,"title":"编辑数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"id":"u:8e74d9aec222","showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"dialog"},"editorSetting":{"behavior":"create"},"icon":"fa fa-plus","id":"u:2f6d4463cb00","label":"新增","level":"primary","type":"button"},{"type":"bulk-actions"},{"align":"right","id":"u:df448b218def","type":"reload"},{"align":"right","id":"u:6046fb0268f2","tpl":"内容","type":"filter-toggler","wrapperComponent":""}],"id":"u:563375ef114c","itemActions":[],"messages":{},"perPage":20,"perPageAvailable":[5,10,20,50,100],"syncLocation":false,"type":"crud"}],"data":{"showFilter":false},"definitions":{"modal-ref-1":{"$$ref":"modal-ref-1","actionType":"dialog","actions":[{"actionType":"cancel","id":"u:c66721e47f3c","label":"取消","type":"button"},{"actionType":"confirm","id":"u:934d39f13398","label":"确定","primary":true,"type":"button"}],"body":[{"actions":[{"label":"提交","level":"primary","onEvent":{"click":{"actions":[{"actionType":"submit","componentId":"u:82cdc3b29cd1"}]}},"type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/pages/quick_save"},"body":[{"hidden":false,"id":"u:a96e5b28e7d4","label":"ID","name":"id","row":0,"type":"input-text","visible":false},{"id":"u:7839a70a374f","label":"","language":"json","name":"schema","options":{"autoIndent":"full","cursorBlinking":"smooth","folding":true,"foldingStrategy":"indentation","minimap":{"enabled":true},"quickSuggestions":{"comments":true,"other":true,"strings":true},"validate":true,"wordWrap":"on"},"row":1,"size":"xxl","type":"editor"}],"dsType":"api","feat":"Edit","id":"u:82cdc3b29cd1","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/pages/detail"},"labelAlign":"top","mode":"flex","resetAfterSubmit":true,"title":"表单","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"id":"u:43f35638761d","resizable":true,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"size":"lg","title":"编辑 JSON","type":"dialog"}},"id":"u:a8979bbc7044","pullRefresh":{"disabled":true},"regions":["body"],"type":"page"}`),
		},
		{
			Name:   "管理员",
			Sign:   "admin_user",
			Schema: json.RawMessage(`{"body":[{"alwaysShowPagination":true,"api":{"method":"get","url":"/system/users"},"bulkActions":[{"actionType":"ajax","api":"","confirmText":"确定要删除？","editorSetting":{"behavior":"bulkDelete"},"icon":"fa fa-trash-o","id":"u:394fe8eeb9b7","label":"删除","level":"danger","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${ids}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/users/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button"}],"columns":[{"id":"u:154e022f601d","label":"ID","name":"id","placeholder":"-","sortable":true,"type":"text"},{"id":"u:315b256293b8","inline":false,"label":"头像","name":"avatar","placeholder":"-","showtype":"image","src":"${avatar}","style":{"fontFamily":"","fontSize":12},"type":"avatar"},{"id":"u:38093dfa5519","label":"姓名","name":"name","placeholder":"-","type":"tpl"},{"id":"u:204471470466","label":"用户名","name":"username","placeholder":"-","type":"tpl"},{"id":"u:924252b24efa","isFixedHeight":false,"isFixedWidth":false,"items":{"className":"my-1","id":"u:a9630e9ed74a","label":"${name}","type":"tag"},"label":"角色","name":"roles","placeholder":"-","style":{"display":"block"},"type":"each"},{"id":"u:428d92bcd810","label":"状态","name":"enabled","placeholder":"-","quickEdit":{"body":[{"falseValue":0,"id":"u:f4152133eb3b","label":"","name":"enabled","offText":"禁用","onText":"启用","option":"","trueValue":1,"type":"switch"}],"id":"u:347a3e35a648","isFixedHeight":false,"isFixedWidth":false,"mode":"inline","saveImmediately":true,"style":{"display":"block","position":"static"},"type":"wrapper"},"type":"switch"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:05ee8b110a15","label":"创建时间","name":"created_at","placeholder":"-","sortable":true,"type":"date"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:a17c366510d7","label":"更新时间","name":"updated_at","placeholder":"-","sortable":true,"type":"date"},{"buttons":[{"actionType":"drawer","drawer":{"actionType":"drawer","actions":[{"actionType":"cancel","id":"u:efb060b9ff4a","label":"取消","type":"button"},{"actionType":"submit","id":"u:3f28bcb7d137","label":"提交","level":"primary","type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/users"},"body":[{"colSize":"1","hidden":false,"id":"u:f019150f0d5d","label":"ID","name":"id","row":0,"type":"input-text","visible":false},{"accept":".jpeg, .jpg, .png, .gif","autoUpload":true,"bos":"default","colSize":"1","hideUploadButton":false,"id":"u:6c121c169712","imageClassName":"r w-full","label":"头像","limit":false,"multiple":false,"name":"avatar","proxy":true,"receiver":{"adaptor":"","data":{"dir":"avatars"},"messages":{},"method":"post","requestAdaptor":"","url":"/upload"},"row":1,"type":"input-image","uploadType":"fileReceptor"},{"clearable":true,"colSize":"1","id":"u:879f7940aa10","label":"姓名","name":"name","required":true,"row":2,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"colSize":"1","id":"u:268e0efe869c","label":"用户名","name":"username","required":true,"row":3,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"colSize":"1","id":"u:63199903f7a6","label":"密码","name":"password","row":4,"showCounter":false,"type":"input-password","validations":{"maxLength":255}},{"id":"u:f429d1cc72a8","label":"确认密码","name":"confirm_password","row":5,"showCounter":false,"type":"input-password","validationErrors":{"equalsField":"两次输入密码不一致"},"validations":{"equalsField":"password"}},{"checkAll":false,"clearable":true,"id":"u:4941684849f2","joinValues":true,"label":"角色","multiple":true,"name":"role_ids","row":6,"searchable":true,"source":{"adaptor":"","messages":{},"method":"get","requestAdaptor":"","url":"/system/users/role_options"},"type":"select"},{"falseValue":0,"id":"u:63ea90c7355d","label":"状态","name":"enabled","offText":"禁用","onText":"启用","option":"","row":7,"trueValue":1,"type":"switch","value":1}],"dsType":"api","feat":"Edit","id":"u:b808323b3b92","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/users/detail"},"labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"}]}},"resetAfterSubmit":true,"title":"新增数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"editorSetting":{"displayName":""},"id":"u:eba8ada34731","resizable":false,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"drawer"},"editorSetting":{"behavior":"update"},"id":"u:29435e16d2df","label":"编辑","level":"link","type":"button"},{"actionType":"","api":{"method":"post","url":"/system/users"},"className":"text-danger","confirmText":"确定要删除？","editorSetting":{"behavior":"delete"},"id":"u:cee85dad465d","label":"删除","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${TRIM(id)}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/users/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button"}],"id":"u:01259e435906","label":"操作","type":"operation"}],"filter":{"actions":[{"actionType":"clear-and-submit","id":"u:cdb68b8d4b45","label":"重置","type":"button"},{"id":"u:253f267c0d95","label":"搜索","level":"primary","type":"submit"}],"body":[{"clearable":true,"id":"u:7fed8045c4d1","label":"名称","name":"name","size":"md","type":"input-text"},{"clearable":true,"id":"u:7fed8045c4d1","label":"用户名","name":"username","size":"md","type":"input-text"}],"feat":"Insert","id":"u:ba2b6c5f28ea","panelClassName":"base-filter","title":"","type":"form"},"filterDefaultVisible":false,"filterTogglable":true,"footerToolbar":[{"type":"statistics"},{"type":"pagination"},{"body":[{"id":"u:8378ce2f122e","multiple":false,"name":"perPage","onEvent":{"change":{"actions":[{"actionType":"reload","componentId":"admin-api.system.admin_users.crud","data":{"perPage":"${event.data.value}"}}]}},"options":[{"label":"20 条/页","value":20},{"label":"30 条/页","value":30},{"label":"50 条/页","value":50},{"label":"100 条/页","value":100},{"label":"200 条/页","value":200}],"overlayPlacement":"top","selectFirst":true,"type":"select"}],"feat":"Insert","id":"u:abf9235be81a","target":"window","type":"form","wrapWithPanel":false}],"headerToolbar":[{"actionType":"drawer","drawer":{"actionType":"drawer","actions":[{"actionType":"cancel","id":"u:8468ef2f62a5","label":"取消","type":"button"},{"actionType":"confirm","id":"u:ae926ea9a3e9","label":"提交","primary":true,"type":"button"}],"body":[{"actions":[{"label":"提交","primary":true,"type":"submit"}],"api":{"method":"post","url":"/system/users"},"autoFocus":true,"body":[{"colSize":"1","hidden":false,"id":"u:7a74f685dc9b","label":"ID","name":"id","row":0,"type":"input-text","visible":false},{"accept":".jpeg, .jpg, .png, .gif","autoUpload":true,"bos":"default","colSize":"1","hideUploadButton":false,"id":"u:1780a8c481c4","imageClassName":"r w-full","label":"头像","limit":false,"multiple":false,"name":"avatar","proxy":true,"receiver":{"adaptor":"","data":{"dir":"avatars"},"messages":{},"method":"post","requestAdaptor":"","url":"/upload"},"row":1,"type":"input-image","uploadType":"fileReceptor"},{"clearable":true,"colSize":"1","id":"u:7bf0c6aa5731","label":"姓名","name":"name","required":true,"row":2,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"colSize":"1","id":"u:de8fafd6a92b","label":"用户名","name":"username","required":true,"row":3,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"colSize":"1","id":"u:5dc8f42edf1d","label":"密码","name":"password","required":true,"row":4,"showCounter":false,"type":"input-password","validations":{"maxLength":255}},{"id":"u:39f0f1a0f986","label":"确认密码","name":"confirm_password","row":5,"showCounter":false,"type":"input-password","validationErrors":{"equalsField":"两次输入密码不一致"},"validations":{"equalsField":"password"}},{"checkAll":false,"clearable":true,"id":"u:68e12beb0fb3","joinValues":true,"label":"角色","multiple":true,"name":"role_ids","row":6,"searchable":true,"source":{"adaptor":"","messages":{},"method":"get","requestAdaptor":"","url":"/system/users/role_options"},"type":"select"},{"falseValue":0,"id":"u:59675f8a4d35","label":"状态","name":"enabled","offText":"禁用","onText":"启用","option":"","row":7,"trueValue":1,"type":"switch","value":1}],"canAccessSuperData":false,"dsType":"api","feat":"Insert","id":"u:d21c51525db6","labelAlign":"left","mode":"normal","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"id":"u:5778edd54fea","resizable":false,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"新增","type":"drawer"},"editorSetting":{"behavior":"create"},"icon":"fa fa-plus","id":"u:2f6d4463cb00","label":"新增","level":"primary","type":"button"},{"type":"bulk-actions"},{"align":"right","id":"u:df448b218def","type":"reload"},{"align":"right","id":"u:6046fb0268f2","tpl":"内容","type":"filter-toggler","wrapperComponent":""}],"id":"u:563375ef114c","itemActions":[],"messages":{},"perPage":20,"perPageAvailable":[5,10,20,50,100],"quickSaveItemApi":{"method":"post","url":"/system/users/quick_save"},"syncLocation":false,"type":"crud"}],"data":{"showFilter":false},"definitions":{},"id":"u:a8979bbc7044","pullRefresh":{"disabled":true},"regions":["body"],"type":"page"}`),
		},
		{
			Name:   "角色",
			Sign:   "admin_role",
			Schema: json.RawMessage(`{"type":"page","regions":["body"],"id":"u:a8979bbc7044","pullRefresh":{"disabled":true},"body":[{"type":"crud","syncLocation":false,"api":{"method":"get","url":"/system/roles"},"bulkActions":[{"type":"button","level":"danger","label":"删除","actionType":"ajax","confirmText":"确定要删除？","api":"","editorSetting":{"behavior":"bulkDelete"},"id":"u:394fe8eeb9b7","icon":"fa fa-trash-o","onEvent":{"click":{"weight":0,"actions":[{"ignoreError":false,"outputVar":"responseResult","actionType":"ajax","options":{},"api":{"url":"/system/roles/delete","method":"post","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${ids}"}}},{"componentId":"u:563375ef114c","groupType":"component","actionType":"reload"}]}}}],"itemActions":[],"headerToolbar":[{"label":"新增","type":"button","actionType":"dialog","level":"primary","editorSetting":{"behavior":"create"},"dialog":{"type":"dialog","title":"新增","body":[{"type":"form","api":{"method":"post","url":"/system/roles"},"body":[{"type":"input-text","name":"name","label":"名称","id":"u:210c457f5e68","clearable":true,"required":true,"validations":{"maxLength":255}},{"type":"input-text","name":"sign","label":"标识","id":"u:9db61bb91791","clearable":true,"required":true,"validations":{"maxLength":255}}],"id":"u:b49730d65248","actions":[{"type":"submit","label":"提交","primary":true}],"feat":"Insert","dsType":"api","labelAlign":"left","mode":"normal","autoFocus":true,"canAccessSuperData":false}],"actionType":"dialog","id":"u:861566ef1498","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:643352191512"},{"type":"button","actionType":"confirm","label":"提交","primary":true,"id":"u:ecfcc2bf52ae"}],"showCloseButton":true,"closeOnOutside":true,"closeOnEsc":true,"showErrorMsg":true,"showLoading":true,"draggable":true},"id":"u:2f6d4463cb00","icon":"fa fa-plus"},{"type":"bulk-actions"},{"type":"reload","id":"u:df448b218def","align":"right"},{"type":"filter-toggler","tpl":"内容","wrapperComponent":"","id":"u:6046fb0268f2","align":"right"}],"columns":[{"name":"id","label":"ID","type":"text","id":"u:154e022f601d","placeholder":"-","sortable":true},{"name":"name","label":"名称","type":"text","id":"u:7172c1ba66c8","placeholder":"-"},{"type":"text","label":"标识","name":"sign","id":"u:acc2fc18bb5d"},{"type":"date","label":"创建时间","name":"created_at","id":"u:05ee8b110a15","placeholder":"-","sortable":true,"format":"YYYY-MM-DD HH:mm:ss"},{"type":"date","label":"更新时间","name":"updated_at","id":"u:a17c366510d7","placeholder":"-","sortable":true,"format":"YYYY-MM-DD HH:mm:ss"},{"type":"operation","label":"操作","buttons":[{"label":"编辑","type":"button","actionType":"dialog","level":"link","editorSetting":{"behavior":"update"},"dialog":{"type":"dialog","title":"编辑","body":[{"type":"form","api":{"method":"post","url":"/system/roles","requestAdaptor":"","adaptor":"","messages":{}},"initApi":{"method":"get","url":"/system/roles/detail","requestAdaptor":"","adaptor":"","messages":{},"data":{"id":"${id}"}},"body":[{"name":"id","label":"ID","type":"input-text","id":"u:1bc881485c42","visible":false},{"name":"name","label":"名称","type":"input-text","id":"u:be7498337250","clearable":true,"required":true,"validations":{"maxLength":255}},{"label":"标识","name":"sign","type":"input-text","id":"u:18e8745dd901","clearable":true,"required":true,"validations":{"maxLength":255}}],"id":"u:a7cba0632cbb","actions":[{"type":"submit","label":"提交","primary":true}],"feat":"Edit","dsType":"api","labelAlign":"left","mode":"normal","autoFocus":true}],"actionType":"dialog","id":"u:b7639c73fc66","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:75c91a63edad"},{"type":"button","actionType":"confirm","label":"确定","primary":true,"id":"u:f8f7f1c670ea"}],"showCloseButton":true,"closeOnOutside":false,"closeOnEsc":false,"showErrorMsg":true,"showLoading":true,"draggable":false},"id":"u:29435e16d2df"},{"type":"button","label":"删除","actionType":"","level":"link","className":"text-danger","confirmText":"确定要删除？","api":{"method":"post","url":"/system/roles"},"editorSetting":{"behavior":"delete"},"id":"u:cee85dad465d","onEvent":{"click":{"weight":0,"actions":[{"ignoreError":false,"outputVar":"responseResult","actionType":"ajax","options":{},"api":{"url":"/system/roles/delete","method":"post","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${TRIM(id)}"}}},{"componentId":"u:563375ef114c","groupType":"component","actionType":"reload"}]}}}],"id":"u:01259e435906"}],"id":"u:563375ef114c","perPageAvailable":[5,10,20,50,100],"messages":{},"filter":{"type":"form","panelClassName":"base-filter","title":"","actions":[{"type":"button","label":"重置","actionType":"clear-and-submit","id":"u:cdb68b8d4b45"},{"type":"submit","label":"搜索","level":"primary","id":"u:253f267c0d95"}],"body":[{"type":"input-text","name":"name","label":"名称","size":"md","id":"u:7fed8045c4d1","clearable":true},{"type":"input-text","name":"sign","label":"标识","size":"md","id":"u:7fed8045c4d1","clearable":true}],"id":"u:ba2b6c5f28ea","feat":"Insert"},"filterDefaultVisible":false,"filterTogglable":true,"alwaysShowPagination":true,"footerToolbar":[{"type":"statistics"},{"type":"pagination"},{"type":"form","wrapWithPanel":false,"body":[{"type":"select","name":"perPage","options":[{"label":"20 条/页","value":20},{"label":"30 条/页","value":30},{"label":"50 条/页","value":50},{"label":"100 条/页","value":100},{"label":"200 条/页","value":200}],"overlayPlacement":"top","onEvent":{"change":{"actions":[{"actionType":"reload","componentId":"admin-api.system.admin_users.crud","data":{"perPage":"${event.data.value}"}}]}},"id":"u:8378ce2f122e","multiple":false,"selectFirst":true}],"target":"window","id":"u:abf9235be81a","feat":"Insert"}],"perPage":20}],"definitions":{},"data":{"showFilter":false}}`),
		},
		{
			Name:   "菜单",
			Sign:   "admin_menu",
			Schema: json.RawMessage(`{"body":[{"alwaysShowPagination":true,"api":{"method":"get","url":"/system/menus"},"bulkActions":[{"actionType":"ajax","api":"","confirmText":"确定要删除？","editorSetting":{"behavior":"bulkDelete"},"icon":"fa fa-trash-o","id":"u:394fe8eeb9b7","label":"删除","level":"danger","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${ids}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/menus/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button"}],"columns":[{"id":"u:154e022f601d","label":"ID","name":"id","placeholder":"-","sortable":true,"type":"text"},{"className":"text-center h-full","id":"u:f6bc40b7d0db","items":[{"body":[{"className":"mr-3 text-xl h-full","icon":"${icon}","id":"u:0bdbe2fa7da7","type":"custom-svg-icon"}],"id":"u:36e21992b2fe","size":"none","type":"wrapper"},{"id":"u:fdfb43c65c06","tpl":"${name}","type":"tpl"}],"justify":"start","label":"名称","name":"name","placeholder":"-","type":"flex"},{"id":"u:acc2fc18bb5d","label":"路径","name":"path","placeholder":"-","type":"text"},{"id":"u:0830ba2402e4","inline":true,"label":"排序","name":"sort","placeholder":"-","tpl":"","type":"tpl","wrapperComponent":""},{"id":"u:850548ae097c","inline":true,"label":"是否显示","name":"visible","placeholder":"-","quickEdit":{"actions":[{"label":"提交","primary":true,"type":"submit"}],"body":[{"falseValue":0,"id":"u:dff829169d82","label":"","name":"visible","option":"","optionType":"default","trueValue":1,"type":"checkbox","value":1}],"feat":"Insert","id":"u:ad890db0044b","isFixedHeight":false,"isFixedWidth":false,"mode":"inline","saveImmediately":true,"style":{"display":"block","fontFamily":"","fontSize":12,"position":"static"},"type":"wrapper"},"tableLayout":"fixed","textOverflow":"default","type":"tpl","width":60},{"id":"u:c07dc72a8da7","label":"首页","name":"is_home","onEvent":{},"option":"","optionType":"default","placeholder":"-","quickEdit":{"actions":[{"label":"提交","primary":true,"type":"submit"}],"body":[{"falseValue":0,"id":"u:4a63cc59fdac","label":"","name":"visible","onEvent":{},"option":"","optionType":"default","trueValue":1,"type":"checkbox","value":1}],"feat":"Insert","id":"u:acd5bc492680","isFixedHeight":false,"isFixedWidth":false,"mode":"inline","saveImmediately":{"api":""},"style":{"display":"block","fontFamily":"","fontSize":12,"position":"static"},"type":"wrapper"},"type":"tpl","width":60},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:05ee8b110a15","label":"创建时间","name":"created_at","placeholder":"-","sortable":true,"type":"date"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:a17c366510d7","label":"更新时间","name":"updated_at","placeholder":"-","sortable":true,"type":"date"},{"buttons":[{"actionType":"dialog","dialog":{"actionType":"dialog","actions":[{"actionType":"cancel","id":"u:75c91a63edad","label":"取消","type":"button"},{"actionType":"confirm","id":"u:f8f7f1c670ea","label":"确定","primary":true,"type":"button"}],"body":[{"actions":[{"label":"提交","primary":true,"type":"submit"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/menus"},"autoFocus":true,"body":[{"id":"u:1bc881485c42","label":"ID","name":"id","type":"input-text","visible":false},{"clearable":true,"id":"u:be7498337250","label":"名称","name":"name","required":true,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:18e8745dd901","label":"标识","name":"sign","required":true,"type":"input-text","validations":{"maxLength":255}}],"dsType":"api","feat":"Edit","id":"u:a7cba0632cbb","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/menus/detail"},"labelAlign":"left","mode":"normal","type":"form"}],"closeOnEsc":false,"closeOnOutside":false,"draggable":false,"id":"u:b7639c73fc66","showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"dialog"},"editorSetting":{"behavior":"update"},"id":"u:29435e16d2df","label":"编辑","level":"link","type":"button"},{"actionType":"","api":{"method":"post","url":"/system/menus"},"className":"text-danger","confirmText":"确定要删除？","editorSetting":{"behavior":"delete"},"id":"u:cee85dad465d","label":"删除","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${TRIM(id)}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/menus/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button"}],"id":"u:01259e435906","label":"操作","type":"operation"}],"draggable":true,"filter":{"actions":[{"actionType":"clear-and-submit","id":"u:cdb68b8d4b45","label":"重置","type":"button"},{"id":"u:253f267c0d95","label":"搜索","level":"primary","type":"submit"}],"body":[{"clearable":true,"id":"u:7fed8045c4d1","label":"名称","name":"name","size":"md","type":"input-text"},{"clearable":true,"id":"u:7fed8045c4d1","label":"标识","name":"sign","size":"md","type":"input-text"}],"feat":"Insert","id":"u:ba2b6c5f28ea","panelClassName":"base-filter","title":"","type":"form"},"filterDefaultVisible":false,"filterTogglable":true,"footerToolbar":[{"type":"statistics"},{"type":"pagination"},{"body":[{"id":"u:8378ce2f122e","multiple":false,"name":"perPage","onEvent":{"change":{"actions":[{"actionType":"reload","componentId":"admin-api.system.admin_users.crud","data":{"perPage":"${event.data.value}"}}]}},"options":[{"label":"20 条/页","value":20},{"label":"30 条/页","value":30},{"label":"50 条/页","value":50},{"label":"100 条/页","value":100},{"label":"200 条/页","value":200}],"overlayPlacement":"top","selectFirst":true,"type":"select"}],"feat":"Insert","id":"u:abf9235be81a","target":"window","type":"form","wrapWithPanel":false}],"headerToolbar":[{"actionType":"drawer","drawer":{"actionType":"drawer","actions":[{"actionType":"cancel","id":"u:be2fa4ce4c2f","label":"取消","type":"button"},{"actionType":"submit","id":"u:067186cd7225","label":"提交","level":"primary","type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/menus"},"body":[{"clearable":true,"id":"u:2b942f0e679d","label":"名称","name":"name","required":true,"row":0,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:5af24fc999ee","label":"路径","name":"path","required":true,"row":1,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"description":"在 <a href=\"https://icones.js.org/\" target=\"_blank\">icones</a> 中寻找可用图标","id":"u:250aa465b533","label":"图标","name":"icon","row":2,"type":"input-text","validations":{"maxLength":255}},{"id":"u:590f1768da53","label":"父级","name":"parentId","row":3,"type":"input-text"},{"description":"数字越大越靠前","id":"u:ceb137b81027","keyboard":true,"label":"排序","max":999999999,"min":0,"name":"sort","required":true,"row":4,"showSteps":true,"step":1,"type":"input-number","value":0},{"description":"是否作为菜单显示","falseValue":0,"id":"u:e565472fd3b7","label":"是否显示","name":"visible","option":"","row":5,"trueValue":1,"type":"switch","value":1},{"animations":{},"body":[],"id":"u:dcac272b5437","isFixedHeight":false,"isFixedWidth":false,"isFreeContainer":true,"row":6,"size":"sm","style":{"position":"relative"},"type":"container","wrapperBody":false},{"body":[{"description":"登录后跳转的页面","falseValue":0,"id":"u:ebef1b751c29","label":"首页","name":"isHome","option":"","row":7,"trueValue":1,"type":"switch","value":0},{"colSize":"1","description":"开启后, 该页面将不展示菜单栏","falseValue":0,"id":"u:ddb46c400a4d","label":"全屏","name":"isFull","option":"","row":7,"trueValue":1,"type":"switch"},{"colSize":"1","description":"开启后页面将缓存，重新打开时不会重新加载","falseValue":0,"id":"u:bd3d015a0700","label":"缓存页面","name":"keepAlive","option":"","row":7,"trueValue":1,"type":"switch"}],"collapsable":true,"collapsed":true,"id":"u:fd0975da5389","row":7,"subFormMode":"","title":"更多设置","type":"fieldset"}],"dsType":"api","feat":"Insert","id":"u:74c3b893eaba","labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"}]}},"resetAfterSubmit":true,"title":"新增数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"hideActions":false,"id":"u:9e2aed5b464f","resizable":false,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"新增","type":"drawer"},"editorSetting":{"behavior":"create"},"icon":"fa fa-plus","id":"u:2f6d4463cb00","label":"新增","level":"primary","type":"button"},{"type":"bulk-actions"},{"align":"right","id":"u:df448b218def","type":"reload"},{"align":"right","id":"u:6046fb0268f2","tpl":"内容","type":"filter-toggler","wrapperComponent":""}],"id":"u:563375ef114c","itemActions":[],"messages":{},"onEvent":{"quickSaveItemSucc":{"actions":[{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}],"weight":0},"saveOrderSucc":{"actions":[{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}],"weight":0}},"perPage":20,"perPageAvailable":[5,10,20,50,100],"quickSaveItemApi":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/menus/quick_save"},"saveOrderApi":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/menus/sort"},"syncLocation":false,"type":"crud"}],"data":{"showFilter":false},"definitions":{},"id":"u:a8979bbc7044","pullRefresh":{"disabled":true},"regions":["body"],"type":"page"}`),
		},
	})
}
