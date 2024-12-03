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
			Schema: json.RawMessage(`{"type":"page","regions":["body"],"id":"u:a8979bbc7044","pullRefresh":{"disabled":true},"body":[{"type":"panel","title":"","body":[{"type":"tpl","tpl":"","wrapperComponent":"","inline":false,"id":"u:921ea986c7c6"},{"id":"u:ebed804f316f","type":"crud2","mode":"table2","dsType":"api","syncLocation":true,"selectable":true,"multiple":true,"primaryField":"id","loadType":"pagination","api":{"url":"/system/pages","method":"get"},"quickSaveItemApi":{"url":"/system/pages","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"filter":{"type":"form","title":"筛选","mode":"inline","columnCount":4,"clearValueOnHidden":true,"behavior":["SimpleQuery"],"body":[{"name":"name","label":"名称","type":"input-text","size":"full","required":false,"behavior":"SimpleQuery","id":"u:e6ad3142299a","clearable":true},{"name":"sign","label":"标识","type":"input-text","size":"full","required":false,"behavior":"SimpleQuery","id":"u:0f93933f2010","clearable":true}],"actions":[{"type":"reset","label":"重置","id":"u:e5e0860c35e7"},{"type":"submit","label":"查询","level":"primary","id":"u:dfddb6d9b845"}],"id":"u:625a8a01302f","feat":"Insert","autoFocus":false,"wrapWithPanel":true,"affixFooter":true,"themeCss":{"headerControlClassName":{}},"visibleOn":"${showFilter}","labelAlign":"left"},"headerToolbar":[{"type":"flex","direction":"row","justify":"flex-start","alignItems":"stretch","style":{"position":"static"},"items":[{"type":"container","align":"left","behavior":["Insert","BulkEdit","BulkDelete"],"body":[{"type":"button","label":"新增","level":"primary","className":"m-r-xs","behavior":"Insert","onEvent":{"click":{"actions":[{"actionType":"dialog","dialog":{"type":"dialog","body":[{"id":"u:077351270b41","type":"form","title":"新增数据","mode":"flex","labelAlign":"top","dsType":"api","feat":"Insert","body":[{"name":"name","label":"名称","row":0,"type":"input-text","id":"u:9f939987f6cb","clearable":true,"required":true,"validations":{"maxLength":255}},{"name":"sign","label":"标识","row":1,"type":"input-text","id":"u:2ef00c8e400d","required":true,"validations":{"maxLength":255},"clearable":true},{"type":"input-sub-form","name":"page","label":"页面结构","form":{"type":"form","className":"h-full","size":"full","title":"","body":{"type":"custom-amis-editor","name":"schema","label":"","mode":"normal","className":"h-full"}},"required":true,"id":"u:a45d19071e36"}],"api":{"url":"/system/pages","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"resetAfterSubmit":true,"actions":[{"type":"button","actionType":"cancel","label":"取消"},{"type":"button","actionType":"submit","label":"提交","level":"primary"}],"onEvent":{"submitSucc":{"actions":[{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f"}]}}}],"title":"新增","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:753bbec4a163"},{"type":"button","actionType":"submit","label":"提交","level":"primary","id":"u:a6781c3c01bd"}],"actionType":"dialog","id":"u:fe8bcc856a05","showCloseButton":true,"closeOnOutside":true,"closeOnEsc":true,"showErrorMsg":true,"showLoading":true,"draggable":true}}]}},"id":"u:be3d0c329604","icon":"fa fa-plus"},{"type":"button","label":"删除","behavior":"BulkDelete","level":"danger","className":"m-r-xs","confirmText":"确认删除选中项?","disabledOn":"${selectedItems != null && selectedItems.length < 1}","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"method":"post","url":"/system/pages/delete","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${ids}"}},"outputVar":"responseResult","options":{}},{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f","args":{},"expression":"${responseResult.responseStatus == 0}"}]}},"id":"u:f6976a761eab","icon":"fa fa-trash-o"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexBasis":"auto","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-start"},"id":"u:d2feb4775db8"},{"type":"container","align":"right","behavior":["FuzzyQuery"],"body":[{"type":"button","label":"筛选","onEvent":{"click":{"actions":[{"componentId":"u:a8979bbc7044","ignoreError":false,"actionType":"setValue","args":{"value":{"showFilter":"${!showFilter}"}}}]}},"id":"u:341a73d9cbf1","icon":"fa fa-search","className":{"text-primary":"${showFilter}","border-primary":"${showFilter}"}},{"type":"button","label":"","onEvent":{"click":{"actions":[{"componentId":"u:ebed804f316f","ignoreError":false,"outputVar":"","actionType":"reload"}]}},"id":"u:405b59ca63aa","icon":"fa fa-refresh","className":"m-l-xs"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-end"},"id":"u:9de66d99d8ea","isFixedHeight":false}],"id":"u:b2da8147240d"}],"footerToolbar":[{"type":"flex","direction":"row","justify":"flex-start","alignItems":"stretch","style":{"position":"static"},"items":[{"type":"container","align":"left","body":[],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-start"},"id":"u:4827bfdfabb9","isFixedHeight":false},{"type":"container","align":"right","body":[{"type":"pagination","behavior":"Pagination","layout":["total","perPage","pager","go"],"perPage":20,"perPageAvailable":[20,50,100,200],"align":"right","id":"u:219881e04744","size":"md"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexBasis":"auto","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-end"},"id":"u:6534f7aa4d15"}],"id":"u:39349b5d361d"}],"columns":[{"type":"tpl","title":"ID","name":"id","id":"u:f5854eff0cfa","placeholder":"-","sorter":true},{"type":"tpl","title":"名称","name":"name","id":"u:315b256293b8"},{"type":"tpl","title":"标识","name":"sign","id":"u:204471470466","placeholder":"-"},{"type":"datetime","format":"YYYY-MM-DD HH:mm:ss","value":1732889273,"name":"created_at","title":"创建时间","id":"u:004dde22d187","placeholder":"-","sorter":true},{"type":"datetime","format":"YYYY-MM-DD HH:mm:ss","value":1732889273,"name":"updated_at","title":"更新时间","id":"u:94bf7ad7c0be","placeholder":"-","sorter":true},{"type":"operation","title":"操作","buttons":[{"type":"button","label":"复制","level":"link","behavior":"Edit","onEvent":{"click":{"actions":[{"ignoreError":false,"outputVar":"responseResult","actionType":"ajax","options":{},"api":{"url":"/system/pages/copy","method":"get","requestAdaptor":"","adaptor":"","messages":{},"data":{"id":"${id}"}}},{"componentId":"u:ebed804f316f","groupType":"component","actionType":"search","args":{},"expression":"${responseResult.responseStatus == 0}"}]}},"id":"u:7cae73c66cd0"},{"type":"button","label":"编辑","level":"link","behavior":"Edit","onEvent":{"click":{"actions":[{"actionType":"dialog","dialog":{"$ref":"modal-ref-1"}}]}},"id":"u:4658404f79cb"},{"type":"button","label":"删除","behavior":"Delete","className":"m-r-xs text-danger","level":"link","confirmText":"确认要删除数据","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"method":"post","url":"/system/pages/delete","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${TRIM(id)}"}},"outputVar":"responseResult","options":{}},{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f","args":{"query":""},"expression":"${responseResult.responseStatus == 0}"}]}},"id":"u:df00563b639d"}],"id":"u:8136b229cca9"}],"editorSetting":{"mock":{"enable":true,"maxDisplayRows":5}}}],"id":"u:93deb9fe8588","affixFooter":false,"actions":[]}],"definitions":{"modal-ref-1":{"type":"dialog","body":[{"id":"u:213c44c60d36","type":"form","title":"编辑数据","mode":"flex","labelAlign":"top","dsType":"api","feat":"Edit","body":[{"name":"id","label":"id","row":0,"type":"input-number","id":"u:5736f77c8961","keyboard":true,"step":1,"visible":false,"hidden":true},{"name":"name","label":"名称","row":1,"type":"input-text","id":"u:841f35203ece","clearable":true,"required":true,"validations":{"maxLength":255}},{"name":"sign","label":"标识","row":2,"type":"input-text","id":"u:030699b39cac","required":true,"validations":{"maxLength":255},"clearable":true},{"type":"input-sub-form","name":"page","label":"页面结构","form":{"type":"form","className":"h-full","size":"full","title":"","body":{"type":"custom-amis-editor","name":"schema","label":"","mode":"normal","className":"h-full"}},"required":true,"id":"u:a45d19071e36"}],"api":{"url":"/system/pages","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"resetAfterSubmit":true,"initApi":{"url":"/system/pages/detail","method":"get","requestAdaptor":"","adaptor":"","messages":{},"data":{"id":"${id}"}},"actions":[{"type":"button","actionType":"cancel","label":"取消"},{"type":"button","actionType":"submit","label":"提交","level":"primary"}],"onEvent":{"submitSucc":{"actions":[{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f"}]}}}],"title":"编辑","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:97cbeb046707"},{"type":"button","actionType":"submit","label":"提交","level":"primary","id":"u:60ae7a87557d"}],"actionType":"dialog","id":"u:61abf5f368fa","showCloseButton":true,"closeOnOutside":true,"closeOnEsc":true,"showErrorMsg":true,"showLoading":true,"draggable":true,"$$originId":"a6e0473c85c2","$$ref":"modal-ref-1"}},"data":{"showFilter":false}}`),
		},
		{
			Name:   "管理员",
			Sign:   "admin_user",
			Schema: json.RawMessage(`{"type":"page","regions":["body"],"id":"u:a8979bbc7044","pullRefresh":{"disabled":true},"body":[{"type":"panel","title":"","body":[{"type":"tpl","tpl":"","wrapperComponent":"","inline":false,"id":"u:921ea986c7c6"},{"id":"u:ebed804f316f","type":"crud2","mode":"table2","dsType":"api","syncLocation":true,"selectable":true,"multiple":true,"primaryField":"id","loadType":"pagination","api":{"url":"/system/users","method":"get"},"quickSaveItemApi":{"url":"/system/users/quickSave","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"filter":{"type":"form","title":"筛选","mode":"inline","columnCount":4,"clearValueOnHidden":true,"behavior":["SimpleQuery"],"body":[{"name":"name","label":"姓名","type":"input-text","size":"full","required":false,"behavior":"SimpleQuery","id":"u:e6ad3142299a","clearable":true},{"name":"username","label":"用户名","type":"input-text","size":"full","required":false,"behavior":"SimpleQuery","id":"u:0f93933f2010","clearable":true}],"actions":[{"type":"reset","label":"重置","id":"u:e5e0860c35e7"},{"type":"submit","label":"查询","level":"primary","id":"u:dfddb6d9b845"}],"id":"u:625a8a01302f","feat":"Insert","autoFocus":false,"wrapWithPanel":true,"affixFooter":true,"themeCss":{"headerControlClassName":{}},"visibleOn":"${showFilter}","labelAlign":"left"},"headerToolbar":[{"type":"flex","direction":"row","justify":"flex-start","alignItems":"stretch","style":{"position":"static"},"items":[{"type":"container","align":"left","behavior":["Insert","BulkEdit","BulkDelete"],"body":[{"type":"button","label":"新增","level":"primary","className":"m-r-xs","behavior":"Insert","onEvent":{"click":{"actions":[{"actionType":"drawer","drawer":{"type":"drawer","body":[{"id":"u:bd649d833dd2","type":"form","title":"新增数据","mode":"flex","labelAlign":"top","dsType":"api","feat":"Insert","body":[{"type":"input-image","label":"头像","name":"avatar","autoUpload":true,"proxy":true,"uploadType":"fileReceptor","imageClassName":"r w-full","colSize":"1","id":"u:c16c3a53e420","accept":".jpeg, .jpg, .png, .gif","multiple":false,"hideUploadButton":false,"bos":"default","receiver":{"url":"/upload","method":"post","requestAdaptor":"","adaptor":"","messages":{},"data":{"dir":"avatars"}},"limit":false,"row":0},{"name":"name","label":"姓名","row":1,"type":"input-text","id":"u:f87eac6abed8","clearable":true,"required":true,"validations":{"maxLength":255},"colSize":"1"},{"name":"username","label":"用户名","row":2,"type":"input-text","id":"u:64e69a508bfc","required":true,"validations":{"maxLength":255},"clearable":true,"colSize":"1"},{"name":"password","label":"密码","row":3,"type":"input-password","id":"u:8bc6f7ac386a","required":true,"validations":{"maxLength":255},"clearable":true,"colSize":"1","showCounter":false},{"type":"input-password","label":"确认密码","name":"confirm_password","row":4,"id":"u:481d95c36bf5","showCounter":false,"validations":{"equalsField":"password"},"validationErrors":{"equalsField":"两次输入密码不一致"}},{"type":"select","label":"角色","name":"roleIds","row":5,"id":"u:81a9bcf03a5d","multiple":true,"clearable":true,"searchable":true,"checkAll":false,"joinValues":true,"source":{"method":"get","url":"/system/users/roleOptions","requestAdaptor":"","adaptor":"","messages":{}}},{"type":"switch","label":"状态","option":"","name":"enabled","falseValue":0,"trueValue":1,"row":6,"id":"u:0e6ff3e3f062","onText":"启用","offText":"禁用","value":1}],"api":{"url":"/system/users","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"resetAfterSubmit":true,"actions":[{"type":"button","actionType":"cancel","label":"取消"},{"type":"button","actionType":"submit","label":"提交","level":"primary"}],"onEvent":{"submitSucc":{"actions":[{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f"}]}}}],"title":"新增","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:f461d97d1668"},{"type":"button","actionType":"submit","label":"提交","level":"primary","id":"u:d49976419230"}],"actionType":"drawer","id":"u:9343086bb00f","showCloseButton":true,"closeOnOutside":true,"closeOnEsc":true,"showErrorMsg":true,"showLoading":true,"draggable":true,"resizable":false}}]}},"id":"u:be3d0c329604","icon":"fa fa-plus"},{"type":"button","label":"删除","behavior":"BulkDelete","level":"danger","className":"m-r-xs","confirmText":"确认删除选中项?","disabledOn":"${selectedItems != null && selectedItems.length < 1}","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"method":"post","url":"/system/users/delete","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${ids}"}}},{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f","args":{},"expression":"${responseResult.responseStatus == 0}"}]}},"id":"u:f6976a761eab","icon":"fa fa-trash-o"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexBasis":"auto","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-start"},"id":"u:d2feb4775db8"},{"type":"container","align":"right","behavior":["FuzzyQuery"],"body":[{"type":"button","label":"筛选","onEvent":{"click":{"actions":[{"componentId":"u:a8979bbc7044","ignoreError":false,"actionType":"setValue","args":{"value":{"showFilter":"${!showFilter}"}}}]}},"id":"u:341a73d9cbf1","icon":"fa fa-search","className":{"text-primary":"${showFilter}","border-primary":"${showFilter}"}},{"type":"button","label":"","onEvent":{"click":{"actions":[{"componentId":"u:ebed804f316f","ignoreError":false,"outputVar":"","actionType":"reload"}]}},"id":"u:405b59ca63aa","icon":"fa fa-refresh","className":"m-l-xs"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-end"},"id":"u:9de66d99d8ea","isFixedHeight":false}],"id":"u:b2da8147240d"}],"footerToolbar":[{"type":"flex","direction":"row","justify":"flex-start","alignItems":"stretch","style":{"position":"static"},"items":[{"type":"container","align":"left","body":[],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-start"},"id":"u:4827bfdfabb9","isFixedHeight":false},{"type":"container","align":"right","body":[{"type":"pagination","behavior":"Pagination","layout":["total","perPage","pager","go"],"perPage":20,"perPageAvailable":[20,50,100,200],"align":"right","id":"u:219881e04744","size":"md"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexBasis":"auto","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-end"},"id":"u:6534f7aa4d15"}],"id":"u:39349b5d361d"}],"columns":[{"type":"tpl","title":"ID","name":"id","id":"u:f5854eff0cfa","placeholder":"-","sorter":true},{"type":"avatar","title":"头像","name":"avatar","id":"u:315b256293b8","placeholder":"-","inline":false,"showtype":"image","style":{"fontFamily":"","fontSize":12},"src":"${avatar}"},{"type":"tpl","title":"姓名","name":"name","id":"u:38093dfa5519","placeholder":"-"},{"type":"tpl","title":"用户名","name":"username","id":"u:204471470466","placeholder":"-"},{"type":"each","title":"角色","name":"roles","id":"u:924252b24efa","placeholder":"-","style":{"display":"block"},"isFixedHeight":false,"isFixedWidth":false,"items":{"label":"${name}","type":"tag","id":"u:a9630e9ed74a","className":"my-1"}},{"type":"switch","title":"状态","name":"enabled","id":"u:428d92bcd810","placeholder":"-","quickEdit":{"mode":"inline","id":"u:347a3e35a648","saveImmediately":true,"type":"wrapper","body":[{"type":"switch","label":"","option":"","name":"enabled","falseValue":0,"trueValue":1,"id":"u:f4152133eb3b","onText":"启用","offText":"禁用"}],"style":{"position":"static","display":"block"},"isFixedHeight":false,"isFixedWidth":false}},{"type":"datetime","format":"YYYY-MM-DD HH:mm:ss","value":1732889273,"name":"created_at","title":"创建时间","id":"u:004dde22d187","placeholder":"-","sorter":true},{"type":"datetime","format":"YYYY-MM-DD HH:mm:ss","value":1732889273,"name":"updated_at","title":"更新时间","id":"u:3cf158682742","placeholder":"-","sorter":true},{"type":"operation","title":"操作","buttons":[{"type":"button","label":"编辑","level":"link","behavior":"Edit","onEvent":{"click":{"actions":[{"actionType":"drawer","drawer":{"$ref":"modal-ref-1"}}]}},"id":"u:7cae73c66cd0"},{"type":"button","label":"删除","behavior":"Delete","className":"m-r-xs text-danger","level":"link","confirmText":"确认要删除数据","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"method":"post","url":"/system/users/delete","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${TRIM(id)}"}},"outputVar":"responseResult","options":{}},{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f","args":{"query":""},"expression":"${responseResult.responseStatus == 0}"}]}},"id":"u:df00563b639d"}],"id":"u:8136b229cca9"}],"editorSetting":{"mock":{"enable":true,"maxDisplayRows":5}},"quickSaveApi":{"url":"","method":"get","requestAdaptor":"","adaptor":"","messages":{}}}],"id":"u:93deb9fe8588","affixFooter":false,"actions":[]}],"definitions":{"modal-ref-1":{"type":"drawer","body":[{"id":"u:b808323b3b92","type":"form","title":"新增数据","mode":"flex","labelAlign":"top","dsType":"api","feat":"Edit","body":[{"type":"input-text","label":"ID","name":"id","row":0,"id":"u:f019150f0d5d","colSize":"1","visible":false,"hidden":false},{"type":"input-image","label":"头像","name":"avatar","autoUpload":true,"proxy":true,"uploadType":"fileReceptor","imageClassName":"r w-full","colSize":"1","id":"u:6c121c169712","accept":".jpeg, .jpg, .png, .gif","multiple":false,"hideUploadButton":false,"bos":"default","receiver":{"url":"/upload","method":"post","requestAdaptor":"","adaptor":"","messages":{},"data":{"dir":"avatars"}},"limit":false,"row":1},{"name":"name","label":"姓名","row":2,"type":"input-text","id":"u:879f7940aa10","clearable":true,"required":true,"validations":{"maxLength":255},"colSize":"1"},{"name":"username","label":"用户名","row":3,"type":"input-text","id":"u:268e0efe869c","required":true,"validations":{"maxLength":255},"clearable":true,"colSize":"1"},{"name":"password","label":"密码","row":4,"type":"input-password","id":"u:63199903f7a6","validations":{"maxLength":255},"clearable":true,"colSize":"1","showCounter":false},{"type":"input-password","label":"确认密码","name":"confirm_password","row":5,"id":"u:f429d1cc72a8","showCounter":false,"validations":{"equalsField":"password"},"validationErrors":{"equalsField":"两次输入密码不一致"}},{"type":"select","label":"角色","name":"roleIds","row":6,"id":"u:4941684849f2","multiple":true,"clearable":true,"searchable":true,"checkAll":false,"joinValues":true,"source":{"method":"get","url":"/system/users/roleOptions","requestAdaptor":"","adaptor":"","messages":{}}},{"type":"switch","label":"状态","option":"","name":"enabled","falseValue":0,"trueValue":1,"row":7,"id":"u:63ea90c7355d","onText":"启用","offText":"禁用","value":1}],"resetAfterSubmit":true,"actions":[{"type":"button","actionType":"cancel","label":"取消"},{"type":"button","actionType":"submit","label":"提交","level":"primary"}],"onEvent":{"submitSucc":{"actions":[{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f"}]}},"initApi":{"method":"get","url":"/system/users/detail","requestAdaptor":"","adaptor":"","messages":{},"data":{"id":"${id}"}},"api":{"url":"/system/users","method":"post","requestAdaptor":"","adaptor":"","messages":{}}}],"title":"编辑","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:efb060b9ff4a"},{"type":"button","actionType":"submit","label":"提交","level":"primary","id":"u:3f28bcb7d137"}],"actionType":"drawer","id":"u:eba8ada34731","showCloseButton":true,"closeOnOutside":true,"closeOnEsc":true,"showErrorMsg":true,"showLoading":true,"draggable":true,"resizable":false,"editorSetting":{"displayName":""},"$$ref":"modal-ref-1"}},"data":{"showFilter":false}}`),
		},
		{
			Name:   "角色",
			Sign:   "admin_role",
			Schema: json.RawMessage(`{"type":"page","regions":["body"],"id":"u:a8979bbc7044","pullRefresh":{"disabled":true},"body":[{"type":"panel","title":"","body":[{"type":"tpl","tpl":"","wrapperComponent":"","inline":false,"id":"u:921ea986c7c6"},{"id":"u:ebed804f316f","type":"crud2","mode":"table2","dsType":"api","syncLocation":true,"selectable":true,"multiple":true,"primaryField":"id","loadType":"pagination","api":{"url":"/system/roles","method":"get"},"quickSaveItemApi":{"url":"/system/roles","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"filter":{"type":"form","title":"筛选","mode":"inline","columnCount":4,"clearValueOnHidden":true,"behavior":["SimpleQuery"],"body":[{"name":"name","label":"名称","type":"input-text","size":"full","required":false,"behavior":"SimpleQuery","id":"u:e6ad3142299a","clearable":true},{"name":"sign","label":"标识","type":"input-text","size":"full","required":false,"behavior":"SimpleQuery","id":"u:0f93933f2010","clearable":true}],"actions":[{"type":"reset","label":"重置","id":"u:e5e0860c35e7"},{"type":"submit","label":"查询","level":"primary","id":"u:dfddb6d9b845"}],"id":"u:625a8a01302f","feat":"Insert","autoFocus":false,"wrapWithPanel":true,"affixFooter":true,"themeCss":{"headerControlClassName":{}},"visibleOn":"${showFilter}","labelAlign":"left"},"headerToolbar":[{"type":"flex","direction":"row","justify":"flex-start","alignItems":"stretch","style":{"position":"static"},"items":[{"type":"container","align":"left","behavior":["Insert","BulkEdit","BulkDelete"],"body":[{"type":"button","label":"新增","level":"primary","className":"m-r-xs","behavior":"Insert","onEvent":{"click":{"actions":[{"actionType":"dialog","dialog":{"type":"dialog","body":[{"id":"u:9417da9f626c","type":"form","title":"新增数据","mode":"flex","labelAlign":"top","dsType":"api","feat":"Insert","body":[{"name":"name","label":"名称","row":0,"type":"input-text","id":"u:98ae11356cf9","clearable":true,"required":true,"validations":{"maxLength":255}},{"name":"sign","label":"标识","row":1,"type":"input-text","id":"u:b3925ee0fb27","required":true,"validations":{"maxLength":255},"clearable":true}],"api":{"url":"/system/roles","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"resetAfterSubmit":true,"actions":[{"type":"button","actionType":"cancel","label":"取消"},{"type":"button","actionType":"submit","label":"提交","level":"primary"}],"onEvent":{"submitSucc":{"actions":[{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f"}]}}}],"title":"新增","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:c57b8853e1fd"},{"type":"button","actionType":"submit","label":"提交","level":"primary","id":"u:01fcb02eeac0"}],"actionType":"dialog","id":"u:288321647479","showCloseButton":true,"closeOnOutside":true,"closeOnEsc":true,"showErrorMsg":true,"showLoading":true,"draggable":true,"hideActions":false,"resizable":true}}]}},"id":"u:be3d0c329604","icon":"fa fa-plus"},{"type":"button","label":"删除","behavior":"BulkDelete","level":"danger","className":"m-r-xs","confirmText":"确认删除选中项?","disabledOn":"${selectedItems != null && selectedItems.length < 1}","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"method":"post","url":"/system/roles/delete","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${ids}"}}},{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f","args":{},"expression":"${responseResult.responseStatus == 0}"}]}},"id":"u:f6976a761eab","icon":"fa fa-trash-o"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexBasis":"auto","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-start"},"id":"u:d2feb4775db8"},{"type":"container","align":"right","behavior":["FuzzyQuery"],"body":[{"type":"button","label":"筛选","onEvent":{"click":{"actions":[{"componentId":"u:a8979bbc7044","ignoreError":false,"actionType":"setValue","args":{"value":{"showFilter":"${!showFilter}"}}}]}},"id":"u:341a73d9cbf1","icon":"fa fa-search","className":{"text-primary":"${showFilter}","border-primary":"${showFilter}"}},{"type":"button","label":"","onEvent":{"click":{"actions":[{"componentId":"u:ebed804f316f","ignoreError":false,"outputVar":"","actionType":"reload"}]}},"id":"u:405b59ca63aa","icon":"fa fa-refresh","className":"m-l-xs"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-end"},"id":"u:9de66d99d8ea","isFixedHeight":false}],"id":"u:b2da8147240d"}],"footerToolbar":[{"type":"flex","direction":"row","justify":"flex-start","alignItems":"stretch","style":{"position":"static"},"items":[{"type":"container","align":"left","body":[],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-start"},"id":"u:4827bfdfabb9","isFixedHeight":false},{"type":"container","align":"right","body":[{"type":"pagination","behavior":"Pagination","layout":["total","perPage","pager","go"],"perPage":20,"perPageAvailable":[20,50,100,200],"align":"right","id":"u:219881e04744","size":"md"}],"wrapperBody":false,"style":{"flexGrow":1,"flex":"1 1 auto","position":"static","display":"flex","flexBasis":"auto","flexDirection":"row","flexWrap":"nowrap","alignItems":"stretch","justifyContent":"flex-end"},"id":"u:6534f7aa4d15"}],"id":"u:39349b5d361d"}],"columns":[{"type":"tpl","title":"ID","name":"id","id":"u:f5854eff0cfa","placeholder":"-","sorter":true},{"type":"tpl","title":"名称","name":"name","id":"u:315b256293b8"},{"type":"tpl","title":"标识","name":"sign","id":"u:204471470466","placeholder":"-"},{"type":"datetime","format":"YYYY-MM-DD HH:mm:ss","value":1732889273,"name":"created_at","title":"创建时间","id":"u:004dde22d187","placeholder":"-","sorter":true},{"type":"datetime","format":"YYYY-MM-DD HH:mm:ss","value":1732889273,"name":"updated_at","title":"更新时间","id":"u:1b631f4e1af2","placeholder":"-","sorter":true},{"type":"operation","title":"操作","buttons":[{"type":"button","label":"编辑","level":"link","behavior":"Edit","onEvent":{"click":{"actions":[{"actionType":"dialog","dialog":{"type":"dialog","body":[{"id":"u:02e5d6864024","type":"form","title":"编辑数据","mode":"flex","labelAlign":"top","dsType":"api","feat":"Edit","body":[{"name":"id","label":"id","row":0,"type":"input-number","id":"u:f374727d1447","keyboard":true,"step":1,"visible":false,"hidden":true},{"name":"name","label":"名称","row":1,"type":"input-text","id":"u:507bfb2384f0","clearable":true,"required":true,"validations":{"maxLength":255}},{"name":"sign","label":"标识","row":2,"type":"input-text","id":"u:2601d73b09f8","required":true,"validations":{"maxLength":255},"clearable":true}],"api":{"url":"/system/roles","method":"post","requestAdaptor":"","adaptor":"","messages":{}},"resetAfterSubmit":true,"initApi":{"url":"/system/roles/detail","method":"get","requestAdaptor":"","adaptor":"","messages":{},"data":{"id":"${id}"}},"actions":[{"type":"button","actionType":"cancel","label":"取消"},{"type":"button","actionType":"submit","label":"提交","level":"primary"}],"onEvent":{"submitSucc":{"actions":[{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f"}]}}}],"title":"编辑","actions":[{"type":"button","actionType":"cancel","label":"取消","id":"u:3b58cea6b4e9"},{"type":"button","actionType":"submit","label":"提交","level":"primary","id":"u:4a2c30ce7331"}],"actionType":"dialog","id":"u:159082e76c7b","showCloseButton":true,"closeOnOutside":true,"closeOnEsc":true,"showErrorMsg":true,"showLoading":true,"draggable":true,"resizable":true}}]}},"id":"u:7cae73c66cd0"},{"type":"button","label":"删除","behavior":"Delete","className":"m-r-xs text-danger","level":"link","confirmText":"确认要删除数据","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"method":"post","url":"/system/roles/delete","requestAdaptor":"","adaptor":"","messages":{},"data":{"ids":"${TRIM(id)}"}},"outputVar":"responseResult","options":{}},{"actionType":"search","groupType":"component","componentId":"u:ebed804f316f","args":{"query":""},"expression":"${responseResult.responseStatus == 0}"}]}},"id":"u:df00563b639d"}],"id":"u:8136b229cca9"}],"editorSetting":{"mock":{"enable":true,"maxDisplayRows":5}}}],"id":"u:93deb9fe8588","affixFooter":false,"actions":[]}],"definitions":{},"data":{"showFilter":false}}`),
		},
	})
}
