package db

import (
	"amis-base/internal/app/admin/models"
	baseModel "amis-base/internal/models"
	"amis-base/internal/pkg/helper"
	"encoding/json"
	"strings"
)

// Seed 填充初始数据
func Seed() {
	// 设置
	go seedSettings()

	// 页面
	go seedPages()

	// 用户
	go seedUsers()

	go func() {
		// 菜单
		seedMenus()

		// 权限
		seedPermissions()

		// 角色
		seedRoles()
	}()
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
	})
}

// 填充角色
func seedRoles() {
	if !isNull(models.AdminRole{}) {
		return
	}

	db.Create(&[]models.AdminRole{
		{
			BaseModel: baseModel.BaseModel{ID: 1},
			Name:      "超级管理员",
			Sign:      "administrator",
		},
		{
			BaseModel: baseModel.BaseModel{ID: 2},
			Name:      "全权限",
			Sign:      "full_permissions",
		},
	})

	var permissions []models.AdminPermission
	db.Model(&models.AdminPermission{}).Find(&permissions)

	withPermissions := make([]map[string]interface{}, 0) // 菜单和权限关联
	for _, permission := range permissions {
		withPermissions = append(withPermissions, map[string]interface{}{
			"admin_role_id":       2,
			"admin_permission_id": permission.ID,
		})
	}

	db.Exec("truncate table admin_role_permission") // 清除关联表数据
	db.Table("admin_role_permission").Create(withPermissions)
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
			Name: "控制台",
			Sign: "dashboard",
			Schema: json.RawMessage(`{"body":[{"className":"mb-1","columns":[{"md":5,"body":[{"body":[{"body":[{"alignItems":"center","className":"h-full","direction":"column","items":[{"src":"/logo.png","type":"image","id":"u:9daad77eb203","enlargeAble":false,"maxScale":200,"minScale":50,"style":{"display":"inline-block"}},{"body":[{"type":"tpl","tpl":"amisBase","wrapperComponent":"","inline":false,"id":"u:78c6859accac"}],"className":"text-3xl mt-9 font-bold","type":"wrapper","id":"u:141d7c7b7994"},{"className":"w-full mt-5","items":[{"actionType":"url","blank":true,"className":"text-lg font-semibold","label":"GitHub","level":"link","type":"action","id":"u:828d7017d768","url":"https://github.com/slowlyo/amis-base"}],"justify":"center","type":"flex","id":"u:1e73654e75eb","style":{"flexWrap":"nowrap","position":"static","flex":"0 0 auto"},"isFixedWidth":false}],"justify":"center","type":"flex","id":"u:ec432b0aaa63"}],"className":"h-full","type":"wrapper","id":"u:580185319605"}],"className":"h-96","md":5,"type":"card","id":"u:c52669cbd811"}],"id":"u:0d41e248c46f"},{"body":[{"items":[{"body":[{"config":{"backgroundColor":"","legend":{"bottom":0,"left":"center"},"series":[{"avoidLabelOverlap":false,"data":[{"name":"Search Engine","value":1048},{"name":"Direct","value":735},{"name":"Email","value":580},{"name":"Union Ads","value":484},{"name":"Video Ads","value":300}],"emphasis":{"label":{"fontSize":"40","fontWeight":"bold","show":true}},"itemStyle":{"borderColor":"#fff","borderRadius":10,"borderWidth":2},"label":{"position":"center","show":false},"labelLine":{"show":false},"name":"Access From","radius":["40%","70%"],"type":"pie"}],"tooltip":{"trigger":"item"}},"height":350,"type":"chart","id":"u:843ea58e52b1"}],"className":"h-96","type":"card","id":"u:e8fb34b03b52"},{"body":[{"html":"<style>\n    .cube-box{ height: 300px; display: flex; align-items: center; justify-content: center; }\n  .cube { width: 100px; height: 100px; position: relative; transform-style: preserve-3d; animation: rotate 10s linear infinite; }\n  .cube:after {\n    content: '';\n    width: 100%;\n    height: 100%;\n    box-shadow: 0 0 50px rgba(0, 0, 0, 0.2);\n    position: absolute;\n    transform-origin: bottom;\n    transform-style: preserve-3d;\n    transform: rotateX(90deg) translateY(50px) translateZ(-50px);\n    background-color: rgba(0, 0, 0, 0.1);\n  }\n  .cube div {\n    background-color: rgba(64, 158, 255, 0.7);\n    position: absolute;\n    width: 100%;\n    height: 100%;\n    border: 1px solid rgb(27, 99, 170);\n    box-shadow: 0 0 60px rgba(64, 158, 255, 0.7);\n  }\n  .cube div:nth-child(1) { transform: translateZ(-50px); animation: shade 10s -5s linear infinite; }\n  .cube div:nth-child(2) { transform: translateZ(50px) rotateY(180deg); animation: shade 10s linear infinite; }\n  .cube div:nth-child(3) { transform-origin: right; transform: translateZ(50px) rotateY(270deg); animation: shade 10s -2.5s linear infinite; }\n  .cube div:nth-child(4) { transform-origin: left; transform: translateZ(50px) rotateY(90deg); animation: shade 10s -7.5s linear infinite; }\n  .cube div:nth-child(5) { transform-origin: bottom; transform: translateZ(50px) rotateX(90deg); background-color: rgba(0, 0, 0, 0.7); }\n  .cube div:nth-child(6) { transform-origin: top; transform: translateZ(50px) rotateX(270deg); }\n\n  @keyframes rotate {\n    0% { transform: rotateX(-15deg) rotateY(0deg); }\n    100% { transform: rotateX(-15deg) rotateY(360deg); }\n  }\n  @keyframes shade { 50% { background-color: rgba(0, 0, 0, 0.7); } }\n</style>\n<div class=\"cube-box\">\n    <div class=\"cube\">\n        <div></div>\n        <div></div>\n        <div></div>\n        <div></div>\n        <div></div>\n        <div></div>\n    </div>\n</div>","type":"tpl","id":"u:91181c523fa5"}],"className":"h-96 ml-4 w-8/12","type":"card","id":"u:651d15809613"}],"type":"flex","id":"u:a87110dceab0"}],"id":"u:04c30664cd87"}],"type":"grid","id":"u:58c7c1c46171"},{"columns":[{"md":8,"body":[{"body":[{"className":"h-96","config":{"backgroundColor":"","grid":{"bottom":30,"left":"7%","right":"3%","top":60},"legend":{"data":["Visits","Bounce Rate"]},"series":[{"areaStyle":[],"data":[182,58,194,51,139,102,85],"name":"Visits","smooth":true,"symbol":"none","type":"line"},{"areaStyle":[],"data":[85,137,65,61,156,116,67],"name":"Bounce Rate","smooth":true,"symbol":"none","type":"line"}],"title":{"text":"Users Behavior"},"tooltip":{"trigger":"axis"},"xAxis":{"boundaryGap":false,"data":["Mon","Tue","Wed","Thu","Fri","Sat","Sun"],"type":"category"},"yAxis":{"type":"value"}},"height":380,"type":"chart","id":"u:abaf3b8f9ef5"}],"className":"clear-card-mb","md":8,"type":"card","id":"u:5b1cf313325b"}],"id":"u:1e31fd4914d7"},{"body":[{"className":"h-full","direction":"column","items":[{"body":[{"html":"<div id=\"clock\" class=\"text-4xl\"></div><div id=\"clock-date\" class=\"mt-5\"></div>","name":"clock","onMount":"const clock = document.getElementById('clock');\nconst tick = () => {\n    clock.innerHTML = (new Date()).toLocaleTimeString();\n    requestAnimationFrame(tick);\n};\ntick();\n\nconst clockDate = document.getElementById('clock-date');\nclockDate.innerHTML = (new Date()).toLocaleDateString();","type":"custom","id":"u:30c23790b6c5"}],"className":"h-full bg-blingbling mb-4","header":{"title":"Clock"},"type":"card","id":"u:465440721700"},{"body":[{"options":{"breaks":true,"html":true},"type":"markdown","value":"` +
				"```" +
				`go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello World\")\n}\n` +
				"```" +
				`","id":"u:38c6c5a50a09"}],"className":"h-full clear-card-mb rounded-md","type":"panel","id":"u:8ba1ac365edc"}],"type":"flex","id":"u:b9cd87d3e6c4"}],"id":"u:539b792de53a"}],"type":"grid","id":"u:d13debfd381c"}],"className":"m:overflow-auto","css":{".bg-blingbling":{"animation":"gradient 60s ease infinite","background":"linear-gradient(to bottom right, #2C3E50, #FD746C, #FF8235, #ffff1c, #92FE9D, #00C9FF, #a044ff, #e73827)","background-repeat":"no-repeat","background-size":"1000% 1000%","color":"#fff"},".bg-blingbling .cxd-Card-title":{"color":"#fff"},".clear-card-mb":{"margin-bottom":"0 !important"},".cxd-Image":{"border":"0"},"@keyframes gradient":["0%{background-position:0% 0%} 50%{background-position:100% 100%} 100%{background-position:0% 0%}"]},"type":"page","id":"u:41e514f138a7","asideResizor":false,"pullRefresh":{"disabled":true}}`),
		},
		{
			Name:   "页面管理",
			Sign:   "admin_page",
			Schema: json.RawMessage(`{"body":[{"alwaysShowPagination":true,"api":{"method":"get","url":"/system/pages"},"bulkActions":[{"actionType":"ajax","api":"","confirmText":"确定要删除？","editorSetting":{"behavior":"bulkDelete"},"icon":"fa fa-trash-o","id":"u:394fe8eeb9b7","label":"删除","level":"danger","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${ids}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/pages/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button"}],"columns":[{"id":"u:154e022f601d","label":"ID","name":"id","placeholder":"-","sortable":true,"type":"text"},{"id":"u:7172c1ba66c8","label":"名称","name":"name","placeholder":"-","type":"text"},{"id":"u:acc2fc18bb5d","label":"标识","name":"sign","type":"text"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:05ee8b110a15","label":"创建时间","name":"created_at","placeholder":"-","sortable":true,"type":"date"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:a17c366510d7","label":"更新时间","name":"updated_at","placeholder":"-","sortable":true,"type":"date"},{"buttons":[{"actionType":"dialog","dialog":{"actionType":"dialog","actions":[{"actionType":"cancel","id":"u:114404b175c8","label":"取消","type":"button"},{"actionType":"submit","id":"u:b86dc1e4380b","label":"提交","level":"primary","type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/pages"},"body":[{"hidden":true,"id":"u:e5e6c2902816","keyboard":true,"label":"id","name":"id","row":0,"step":1,"type":"input-number","visible":false},{"clearable":true,"id":"u:3f288cda43d2","label":"名称","name":"name","required":true,"row":1,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:46707e48fcd8","label":"标识","name":"sign","required":true,"row":2,"type":"input-text","validations":{"maxLength":255}},{"form":{"body":{"className":"h-full","label":"","mode":"normal","name":"schema","type":"custom-amis-editor"},"className":"h-full","size":"full","title":"","type":"form"},"id":"u:624be988419b","label":"页面结构","name":"page","required":true,"type":"input-sub-form"}],"dsType":"api","feat":"Edit","id":"u:7c52686cc8e9","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/pages/detail"},"labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"}]}},"resetAfterSubmit":true,"title":"编辑数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"id":"u:1cc4b9e5e697","showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"dialog"},"editorSetting":{"behavior":"update"},"id":"u:29435e16d2df","label":"编辑","level":"link","type":"button"},{"behavior":"Edit","id":"u:7cae73c66cd0","label":"复制","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/pages/copy"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","args":{},"componentId":"u:563375ef114c","expression":"${responseResult.responseStatus == 0}","groupType":"component"}]}},"type":"button"},{"id":"u:97b94e051b85","label":"JSON","level":"link","onEvent":{"click":{"actions":[{"actionType":"dialog","dialog":{"$ref":"modal-ref-1"},"ignoreError":false}]}},"type":"button"},{"actionType":"","api":{"method":"post","url":"/system/pages"},"className":"text-danger","confirmText":"确定要删除？","editorSetting":{"behavior":"delete"},"id":"u:cee85dad465d","label":"删除","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${TRIM(id)}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/pages/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button"}],"id":"u:01259e435906","label":"操作","placeholder":"-","type":"operation"}],"filter":{"actions":[{"actionType":"clear-and-submit","id":"u:cdb68b8d4b45","label":"重置","type":"button"},{"id":"u:253f267c0d95","label":"搜索","level":"primary","type":"submit"}],"body":[{"clearable":true,"id":"u:7fed8045c4d1","label":"名称","name":"name","size":"md","type":"input-text"},{"clearable":true,"id":"u:7fed8045c4d1","label":"标识","name":"sign","size":"md","type":"input-text"}],"feat":"Insert","id":"u:ba2b6c5f28ea","panelClassName":"base-filter","title":"","type":"form"},"filterDefaultVisible":false,"filterTogglable":true,"footerToolbar":[{"type":"statistics"},{"type":"pagination"},{"body":[{"id":"u:8378ce2f122e","multiple":false,"name":"perPage","onEvent":{"change":{"actions":[{"actionType":"reload","componentId":"admin-api.system.admin_users.crud","data":{"perPage":"${event.data.value}"}}]}},"options":[{"label":"20 条/页","value":20},{"label":"30 条/页","value":30},{"label":"50 条/页","value":50},{"label":"100 条/页","value":100},{"label":"200 条/页","value":200}],"overlayPlacement":"top","selectFirst":true,"type":"select"}],"feat":"Insert","id":"u:abf9235be81a","target":"window","type":"form","wrapWithPanel":false}],"headerToolbar":[{"actionType":"dialog","dialog":{"actionType":"dialog","actions":[{"actionType":"cancel","id":"u:4b83a945c64a","label":"取消","type":"button"},{"actionType":"submit","id":"u:947858a0fae4","label":"提交","level":"primary","type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/pages"},"body":[{"hidden":true,"id":"u:1248c8850238","keyboard":true,"label":"id","name":"id","row":0,"step":1,"type":"input-number","visible":false},{"clearable":true,"id":"u:7541fa023716","label":"名称","name":"name","required":true,"row":1,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:00fba10e51a0","label":"标识","name":"sign","required":true,"row":2,"type":"input-text","validations":{"maxLength":255}},{"form":{"body":{"className":"h-full","label":"","mode":"normal","name":"schema","type":"custom-amis-editor"},"className":"h-full","size":"full","title":"","type":"form"},"id":"u:954dc473afaf","label":"页面结构","name":"page","required":true,"type":"input-sub-form"}],"dsType":"api","feat":"Edit","id":"u:0269868efcdf","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/pages/detail"},"labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"}]}},"resetAfterSubmit":true,"title":"编辑数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"id":"u:8e74d9aec222","showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"dialog"},"editorSetting":{"behavior":"create"},"icon":"fa fa-plus","id":"u:2f6d4463cb00","label":"新增","level":"primary","type":"button"},{"type":"bulk-actions"},{"align":"right","id":"u:df448b218def","type":"reload"},{"align":"right","id":"u:6046fb0268f2","tpl":"内容","type":"filter-toggler","wrapperComponent":""}],"id":"u:563375ef114c","itemActions":[],"messages":{},"perPage":20,"perPageAvailable":[5,10,20,50,100],"syncLocation":false,"type":"crud"}],"data":{"showFilter":false},"definitions":{"modal-ref-1":{"$$ref":"modal-ref-1","actionType":"dialog","actions":[{"actionType":"cancel","id":"u:c66721e47f3c","label":"取消","type":"button"},{"actionType":"confirm","id":"u:934d39f13398","label":"确定","primary":true,"type":"button"}],"body":[{"actions":[{"label":"提交","level":"primary","onEvent":{"click":{"actions":[{"actionType":"submit","componentId":"u:82cdc3b29cd1"}]}},"type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/pages/quick_save"},"body":[{"hidden":false,"id":"u:a96e5b28e7d4","label":"ID","name":"id","row":0,"type":"input-text","visible":false},{"id":"u:7839a70a374f","label":"","language":"json","name":"schema","options":{"autoIndent":"full","cursorBlinking":"smooth","folding":true,"foldingStrategy":"indentation","minimap":{"enabled":true},"quickSuggestions":{"comments":true,"other":true,"strings":true},"validate":true,"wordWrap":"on"},"row":1,"size":"xxl","type":"editor"}],"dsType":"api","feat":"Edit","id":"u:82cdc3b29cd1","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/pages/detail"},"labelAlign":"top","mode":"flex","resetAfterSubmit":true,"title":"表单","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"id":"u:43f35638761d","resizable":true,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"size":"lg","title":"编辑 JSON","type":"dialog"}},"id":"u:a8979bbc7044","pullRefresh":{"disabled":true},"regions":["body"],"type":"page"}`),
		},
		{
			Name:   "管理员",
			Sign:   "admin_user",
			Schema: json.RawMessage(`{"body":[{"alwaysShowPagination":true,"api":{"method":"get","url":"/system/users"},"bulkActions":[{"actionType":"ajax","api":"","confirmText":"确定要删除？","editorSetting":{"behavior":"bulkDelete"},"icon":"fa fa-trash-o","id":"u:394fe8eeb9b7","label":"删除","level":"danger","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${ids}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/users/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_user.bulk_delete\"]}"}],"columns":[{"id":"u:154e022f601d","label":"ID","name":"id","placeholder":"-","sortable":true,"type":"text"},{"id":"u:315b256293b8","inline":false,"label":"头像","name":"avatar","placeholder":"-","showtype":"image","src":"${avatar}","style":{"fontFamily":"","fontSize":12},"type":"avatar"},{"id":"u:38093dfa5519","label":"姓名","name":"name","placeholder":"-","type":"tpl"},{"id":"u:204471470466","label":"用户名","name":"username","placeholder":"-","type":"tpl"},{"id":"u:924252b24efa","isFixedHeight":false,"isFixedWidth":false,"items":{"className":"my-1","id":"u:a9630e9ed74a","label":"${name}","type":"tag"},"label":"角色","name":"roles","placeholder":"-","style":{"display":"block"},"type":"each"},{"id":"u:428d92bcd810","label":"状态","name":"enabled","placeholder":"-","quickEdit":{"body":[{"falseValue":0,"id":"u:f4152133eb3b","label":"","name":"enabled","offText":"禁用","onText":"启用","option":"","trueValue":1,"type":"switch","disabledOn":"${!isSuperAdmin && _permissions[\"system.admin_user.update\"] != 1}"}],"id":"u:347a3e35a648","isFixedHeight":false,"isFixedWidth":false,"mode":"inline","saveImmediately":true,"style":{"display":"block","position":"static"},"type":"wrapper"},"type":"switch"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:05ee8b110a15","label":"创建时间","name":"created_at","placeholder":"-","sortable":true,"type":"date"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:a17c366510d7","label":"更新时间","name":"updated_at","placeholder":"-","sortable":true,"type":"date"},{"buttons":[{"actionType":"drawer","drawer":{"actionType":"drawer","actions":[{"actionType":"cancel","id":"u:efb060b9ff4a","label":"取消","type":"button"},{"actionType":"submit","id":"u:3f28bcb7d137","label":"提交","level":"primary","type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/users"},"body":[{"colSize":"1","hidden":false,"id":"u:f019150f0d5d","label":"ID","name":"id","row":0,"type":"input-text","visible":false},{"accept":".jpeg, .jpg, .png, .gif","autoUpload":true,"bos":"default","colSize":"1","hideUploadButton":false,"id":"u:6c121c169712","imageClassName":"r w-full","label":"头像","limit":false,"multiple":false,"name":"avatar","proxy":true,"receiver":{"adaptor":"","data":{"dir":"avatars"},"messages":{},"method":"post","requestAdaptor":"","url":"/upload"},"row":1,"type":"input-image","uploadType":"fileReceptor"},{"clearable":true,"colSize":"1","id":"u:879f7940aa10","label":"姓名","name":"name","required":true,"row":2,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"colSize":"1","id":"u:268e0efe869c","label":"用户名","name":"username","required":true,"row":3,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"colSize":"1","id":"u:63199903f7a6","label":"密码","name":"password","row":4,"showCounter":false,"type":"input-password","validations":{"maxLength":255}},{"id":"u:f429d1cc72a8","label":"确认密码","name":"confirm_password","row":5,"showCounter":false,"type":"input-password","validationErrors":{"equalsField":"两次输入密码不一致"},"validations":{"equalsField":"password"}},{"checkAll":false,"clearable":true,"id":"u:4941684849f2","joinValues":true,"label":"角色","multiple":true,"name":"role_ids","row":6,"searchable":true,"source":{"adaptor":"","messages":{},"method":"get","requestAdaptor":"","url":"/system/users/role_options"},"type":"select"},{"falseValue":0,"id":"u:63ea90c7355d","label":"状态","name":"enabled","offText":"禁用","onText":"启用","option":"","row":7,"trueValue":1,"type":"switch","value":1}],"dsType":"api","feat":"Edit","id":"u:b808323b3b92","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/users/detail"},"labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"}]}},"resetAfterSubmit":true,"title":"新增数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"editorSetting":{"displayName":""},"id":"u:eba8ada34731","resizable":false,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"drawer"},"editorSetting":{"behavior":"update"},"id":"u:29435e16d2df","label":"编辑","level":"link","type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_user.update\"]}"},{"actionType":"","api":{"method":"post","url":"/system/users"},"className":"text-danger","confirmText":"确定要删除？","editorSetting":{"behavior":"delete"},"id":"u:cee85dad465d","label":"删除","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${TRIM(id)}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/users/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_user.delete\"]}"}],"id":"u:01259e435906","label":"操作","placeholder":"-","type":"operation"}],"filter":{"actions":[{"actionType":"clear-and-submit","id":"u:cdb68b8d4b45","label":"重置","type":"button"},{"id":"u:253f267c0d95","label":"搜索","level":"primary","type":"submit"}],"body":[{"clearable":true,"id":"u:7fed8045c4d1","label":"名称","name":"name","size":"md","type":"input-text"},{"clearable":true,"id":"u:7fed8045c4d1","label":"用户名","name":"username","size":"md","type":"input-text"}],"feat":"Insert","id":"u:ba2b6c5f28ea","panelClassName":"base-filter","title":"","type":"form"},"filterDefaultVisible":false,"filterTogglable":true,"footerToolbar":[{"type":"statistics"},{"type":"pagination"},{"body":[{"id":"u:8378ce2f122e","multiple":false,"name":"perPage","onEvent":{"change":{"actions":[{"actionType":"reload","componentId":"admin-api.system.admin_users.crud","data":{"perPage":"${event.data.value}"}}]}},"options":[{"label":"20 条/页","value":20},{"label":"30 条/页","value":30},{"label":"50 条/页","value":50},{"label":"100 条/页","value":100},{"label":"200 条/页","value":200}],"overlayPlacement":"top","selectFirst":true,"type":"select"}],"feat":"Insert","id":"u:abf9235be81a","target":"window","type":"form","wrapWithPanel":false}],"headerToolbar":[{"actionType":"drawer","drawer":{"actionType":"drawer","actions":[{"actionType":"cancel","id":"u:8468ef2f62a5","label":"取消","type":"button"},{"actionType":"confirm","id":"u:ae926ea9a3e9","label":"提交","primary":true,"type":"button"}],"body":[{"actions":[{"label":"提交","primary":true,"type":"submit"}],"api":{"method":"post","url":"/system/users"},"autoFocus":true,"body":[{"colSize":"1","hidden":false,"id":"u:7a74f685dc9b","label":"ID","name":"id","row":0,"type":"input-text","visible":false},{"accept":".jpeg, .jpg, .png, .gif","autoUpload":true,"bos":"default","colSize":"1","hideUploadButton":false,"id":"u:1780a8c481c4","imageClassName":"r w-full","label":"头像","limit":false,"multiple":false,"name":"avatar","proxy":true,"receiver":{"adaptor":"","data":{"dir":"avatars"},"messages":{},"method":"post","requestAdaptor":"","url":"/upload"},"row":1,"type":"input-image","uploadType":"fileReceptor"},{"clearable":true,"colSize":"1","id":"u:7bf0c6aa5731","label":"姓名","name":"name","required":true,"row":2,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"colSize":"1","id":"u:de8fafd6a92b","label":"用户名","name":"username","required":true,"row":3,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"colSize":"1","id":"u:5dc8f42edf1d","label":"密码","name":"password","required":true,"row":4,"showCounter":false,"type":"input-password","validations":{"maxLength":255}},{"id":"u:39f0f1a0f986","label":"确认密码","name":"confirm_password","row":5,"showCounter":false,"type":"input-password","validationErrors":{"equalsField":"两次输入密码不一致"},"validations":{"equalsField":"password"}},{"checkAll":false,"clearable":true,"id":"u:68e12beb0fb3","joinValues":true,"label":"角色","multiple":true,"name":"role_ids","row":6,"searchable":true,"source":{"adaptor":"","messages":{},"method":"get","requestAdaptor":"","url":"/system/users/role_options"},"type":"select"},{"falseValue":0,"id":"u:59675f8a4d35","label":"状态","name":"enabled","offText":"禁用","onText":"启用","option":"","row":7,"trueValue":1,"type":"switch","value":1}],"canAccessSuperData":false,"dsType":"api","feat":"Insert","id":"u:d21c51525db6","labelAlign":"left","mode":"normal","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"id":"u:5778edd54fea","resizable":false,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"新增","type":"drawer"},"editorSetting":{"behavior":"create"},"icon":"fa fa-plus","id":"u:2f6d4463cb00","label":"新增","level":"primary","type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_user.create\"]}"},{"type":"bulk-actions"},{"align":"right","id":"u:df448b218def","type":"reload"},{"align":"right","id":"u:6046fb0268f2","tpl":"内容","type":"filter-toggler","wrapperComponent":""}],"id":"u:563375ef114c","itemActions":[],"messages":{},"perPage":20,"perPageAvailable":[5,10,20,50,100],"quickSaveItemApi":{"method":"post","url":"/system/users/quick_save"},"syncLocation":false,"type":"crud"}],"data":{"showFilter":false},"definitions":{},"id":"u:a8979bbc7044","initApi":{"method":"get","url":"/permissions"},"pullRefresh":{"disabled":true},"regions":["body"],"type":"page"}`),
		},
		{
			Name:   "角色",
			Sign:   "admin_role",
			Schema: json.RawMessage(`{"body":[{"alwaysShowPagination":true,"api":{"method":"get","url":"/system/roles"},"bulkActions":[{"actionType":"ajax","api":"","confirmText":"确定要删除？","editorSetting":{"behavior":"bulkDelete"},"icon":"fa fa-trash-o","id":"u:394fe8eeb9b7","label":"删除","level":"danger","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${ids}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/roles/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_role.bulk_delete\"]}"}],"columns":[{"id":"u:154e022f601d","label":"ID","name":"id","placeholder":"-","sortable":true,"type":"text"},{"id":"u:7172c1ba66c8","label":"名称","name":"name","placeholder":"-","type":"text"},{"id":"u:acc2fc18bb5d","label":"标识","name":"sign","type":"text"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:05ee8b110a15","label":"创建时间","name":"created_at","placeholder":"-","sortable":true,"type":"date"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:a17c366510d7","label":"更新时间","name":"updated_at","placeholder":"-","sortable":true,"type":"date"},{"buttons":[{"actionType":"drawer","drawer":{"body":[{"api":"/system/roles/permissions","body":[{"cascade":true,"className":"h-full b-none","extractValue":true,"heightAuto":true,"inputClassName":"h-full tree-full","joinValues":false,"label":"","labelField":"name","multiple":true,"name":"permissions","searchable":true,"size":"full","source":"/system/permissions/parent_options?append_none_option=false","type":"input-tree","valueField":"id"}],"data":{"id":"${id}"},"initApi":"/system/roles/permissions?id=${id}","mode":"normal","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"resizable":true,"title":"设置权限","type":"drawer"},"id":"u:943e80e0bab7","label":"设置权限","level":"link","type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_role.edit_permission\"]}"},{"actionType":"dialog","dialog":{"actionType":"dialog","actions":[{"actionType":"cancel","id":"u:75c91a63edad","label":"取消","type":"button"},{"actionType":"confirm","id":"u:f8f7f1c670ea","label":"确定","primary":true,"type":"button"}],"body":[{"actions":[{"label":"提交","primary":true,"type":"submit"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/roles"},"autoFocus":true,"body":[{"id":"u:1bc881485c42","label":"ID","name":"id","type":"input-text","visible":false},{"clearable":true,"id":"u:be7498337250","label":"名称","name":"name","required":true,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:18e8745dd901","label":"标识","name":"sign","required":true,"type":"input-text","validations":{"maxLength":255}}],"dsType":"api","feat":"Edit","id":"u:a7cba0632cbb","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/roles/detail"},"labelAlign":"left","mode":"normal","type":"form"}],"closeOnEsc":false,"closeOnOutside":false,"draggable":false,"id":"u:b7639c73fc66","showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"dialog"},"editorSetting":{"behavior":"update"},"id":"u:29435e16d2df","label":"编辑","level":"link","type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_role.update\"]}"},{"actionType":"","api":{"method":"post","url":"/system/roles"},"className":"text-danger","confirmText":"确定要删除？","editorSetting":{"behavior":"delete"},"id":"u:cee85dad465d","label":"删除","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${TRIM(id)}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/roles/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"}],"weight":0}},"type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_role.delete\"]}"}],"id":"u:01259e435906","label":"操作","type":"operation"}],"filter":{"actions":[{"actionType":"clear-and-submit","id":"u:cdb68b8d4b45","label":"重置","type":"button"},{"id":"u:253f267c0d95","label":"搜索","level":"primary","type":"submit"}],"body":[{"clearable":true,"id":"u:7fed8045c4d1","label":"名称","name":"name","size":"md","type":"input-text"},{"clearable":true,"id":"u:7fed8045c4d1","label":"标识","name":"sign","size":"md","type":"input-text"}],"feat":"Insert","id":"u:ba2b6c5f28ea","panelClassName":"base-filter","title":"","type":"form"},"filterDefaultVisible":false,"filterTogglable":true,"footerToolbar":[{"type":"statistics"},{"type":"pagination"},{"body":[{"id":"u:8378ce2f122e","multiple":false,"name":"perPage","onEvent":{"change":{"actions":[{"actionType":"reload","componentId":"admin-api.system.admin_users.crud","data":{"perPage":"${event.data.value}"}}]}},"options":[{"label":"20 条/页","value":20},{"label":"30 条/页","value":30},{"label":"50 条/页","value":50},{"label":"100 条/页","value":100},{"label":"200 条/页","value":200}],"overlayPlacement":"top","selectFirst":true,"type":"select"}],"feat":"Insert","id":"u:abf9235be81a","target":"window","type":"form","wrapWithPanel":false}],"headerToolbar":[{"actionType":"dialog","dialog":{"actionType":"dialog","actions":[{"actionType":"cancel","id":"u:643352191512","label":"取消","type":"button"},{"actionType":"confirm","id":"u:ecfcc2bf52ae","label":"提交","primary":true,"type":"button"}],"body":[{"actions":[{"label":"提交","primary":true,"type":"submit"}],"api":{"method":"post","url":"/system/roles"},"autoFocus":true,"body":[{"clearable":true,"id":"u:210c457f5e68","label":"名称","name":"name","required":true,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:9db61bb91791","label":"标识","name":"sign","required":true,"type":"input-text","validations":{"maxLength":255}}],"canAccessSuperData":false,"dsType":"api","feat":"Insert","id":"u:b49730d65248","labelAlign":"left","mode":"normal","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"id":"u:861566ef1498","showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"新增","type":"dialog"},"editorSetting":{"behavior":"create"},"icon":"fa fa-plus","id":"u:2f6d4463cb00","label":"新增","level":"primary","type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_role.create\"]}"},{"type":"bulk-actions"},{"align":"right","id":"u:df448b218def","type":"reload"},{"align":"right","id":"u:6046fb0268f2","tpl":"内容","type":"filter-toggler","wrapperComponent":""}],"id":"u:563375ef114c","itemActions":[],"messages":{},"perPage":20,"perPageAvailable":[5,10,20,50,100],"syncLocation":false,"type":"crud"}],"data":{"showFilter":false},"definitions":{},"id":"u:a8979bbc7044","initApi":{"method":"get","url":"/permissions"},"pullRefresh":{"disabled":true},"regions":["body"],"type":"page"}`),
		},
		{
			Name:   "菜单",
			Sign:   "admin_menu",
			Schema: json.RawMessage(`{"body":[{"alwaysShowPagination":true,"api":{"method":"get","url":"/system/menus"},"bulkActions":[{"actionType":"ajax","api":"","confirmText":"确定要删除？","editorSetting":{"behavior":"bulkDelete"},"icon":"fa fa-trash-o","id":"u:394fe8eeb9b7","label":"删除","level":"danger","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${ids}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/menus/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"},{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}],"weight":0}},"type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_menu.bulk_delete\"]}"}],"columns":[{"id":"u:154e022f601d","label":"ID","name":"id","placeholder":"-","sortable":true,"type":"text"},{"className":"text-center h-full","id":"u:f6bc40b7d0db","items":[{"body":[{"className":"mr-3 text-xl h-full","icon":"${icon}","id":"u:0bdbe2fa7da7","type":"custom-svg-icon"}],"id":"u:36e21992b2fe","size":"none","type":"wrapper"},{"id":"u:fdfb43c65c06","tpl":"${name}","type":"tpl"}],"justify":"start","label":"名称","name":"name","placeholder":"-","type":"flex"},{"id":"u:acc2fc18bb5d","label":"路径","name":"path","placeholder":"-","type":"text"},{"id":"u:0830ba2402e4","inline":true,"label":"排序","name":"sort","placeholder":"-","tpl":"","type":"tpl","wrapperComponent":""},{"id":"u:38fcde7f85e5","inline":true,"label":"页面","name":"page.name","placeholder":"-","type":"tpl","wrapperComponent":""},{"id":"u:850548ae097c","inline":true,"label":"是否显示","name":"visible","placeholder":"-","quickEdit":{"actions":[{"label":"提交","primary":true,"type":"submit"}],"body":[{"falseValue":0,"id":"u:dff829169d82","label":"","name":"visible","option":"","optionType":"default","trueValue":1,"type":"checkbox","value":1,"disabledOn":"${!isSuperAdmin && _permissions[\"system.admin_menu.update\"] != 1}"}],"feat":"Insert","id":"u:ad890db0044b","isFixedHeight":false,"isFixedWidth":false,"mode":"inline","saveImmediately":true,"style":{"display":"block","fontFamily":"","fontSize":12,"position":"static"},"type":"wrapper"},"tableLayout":"fixed","textOverflow":"default","type":"tpl","width":60},{"id":"u:c07dc72a8da7","label":"首页","name":"is_home","onEvent":{},"option":"","optionType":"default","placeholder":"-","quickEdit":{"actions":[{"label":"提交","primary":true,"type":"submit"}],"body":[{"falseValue":0,"id":"u:4a63cc59fdac","label":"","name":"is_home","onEvent":{},"option":"","optionType":"default","trueValue":1,"type":"checkbox","value":1,"disabledOn":"${!isSuperAdmin && _permissions[\"system.admin_menu.update\"] != 1}"}],"feat":"Insert","id":"u:acd5bc492680","isFixedHeight":false,"isFixedWidth":false,"mode":"inline","saveImmediately":{"api":""},"style":{"display":"block","fontFamily":"","fontSize":12,"position":"static"},"type":"wrapper"},"type":"tpl","width":""},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:05ee8b110a15","label":"创建时间","name":"created_at","placeholder":"-","sortable":true,"type":"date"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:a17c366510d7","label":"更新时间","name":"updated_at","placeholder":"-","sortable":true,"type":"date"},{"buttons":[{"actionType":"drawer","drawer":{"actionType":"drawer","actions":[{"actionType":"cancel","id":"u:9a6baf18c2c8","label":"取消","type":"button"},{"actionType":"confirm","id":"u:ce271dcff20b","label":"确定","primary":true,"type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/menus"},"autoFocus":true,"body":[{"colSize":"1","id":"u:b8e7e113f919","label":"ID","name":"id","row":0,"type":"input-text","visible":false},{"clearable":true,"colSize":"1","id":"u:2b942f0e679d","label":"名称","name":"name","required":true,"row":1,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:5af24fc999ee","label":"路径","name":"path","required":true,"row":2,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"description":"在 <a href=\"https://icones.js.org/\" target=\"_blank\">icones</a> 中寻找可用图标","id":"u:250aa465b533","label":"图标","name":"icon","row":3,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"enableNodePath":false,"id":"u:590f1768da53","initiallyOpen":true,"label":"父级","labelField":"name","multiple":false,"name":"parent_id","row":4,"searchable":true,"showIcon":false,"source":"/system/menus/parent_options","type":"tree-select","valueField":"id"},{"description":"数字越大越靠前","id":"u:ceb137b81027","keyboard":true,"label":"排序","max":999999999,"min":0,"name":"sort","required":true,"row":5,"showSteps":true,"step":1,"type":"input-number","value":0},{"clearable":true,"id":"u:9341a4d17c70","label":"页面","multiple":false,"name":"page_sign","row":6,"searchable":true,"source":{"adaptor":"","messages":{},"method":"get","requestAdaptor":"","url":"/system/menus/page_options"},"type":"select"},{"description":"是否作为菜单显示","falseValue":0,"id":"u:e565472fd3b7","label":"是否显示","name":"visible","option":"","row":7,"trueValue":1,"type":"switch","value":1},{"animations":{},"body":[],"id":"u:dcac272b5437","isFixedHeight":false,"isFixedWidth":false,"isFreeContainer":true,"row":8,"size":"sm","style":{"position":"relative"},"type":"container","wrapperBody":false},{"body":[{"description":"登录后跳转的页面","falseValue":0,"id":"u:ebef1b751c29","label":"首页","name":"is_home","option":"","row":7,"trueValue":1,"type":"switch","value":0},{"colSize":"1","description":"开启后, 该页面将不展示菜单栏","falseValue":0,"id":"u:ddb46c400a4d","label":"全屏","name":"is_full","option":"","row":7,"trueValue":1,"type":"switch"},{"colSize":"1","description":"开启后页面将缓存，重新打开时不会重新加载","falseValue":0,"id":"u:bd3d015a0700","label":"缓存页面","name":"keep_alive","option":"","row":7,"trueValue":1,"type":"switch"}],"collapsable":true,"collapsed":true,"id":"u:fd0975da5389","row":9,"subFormMode":"","title":"更多设置","type":"fieldset"}],"dsType":"api","feat":"Edit","id":"u:74c3b893eaba","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/menus/detail"},"labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"},{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}]}},"resetAfterSubmit":true,"title":"新增数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":false,"id":"u:cc540ca818a8","resizable":false,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"drawer"},"editorSetting":{"behavior":"update"},"id":"u:29435e16d2df","label":"编辑","level":"link","type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_menu.update\"]}"},{"actionType":"","api":{"method":"post","url":"/system/menus"},"className":"text-danger","confirmText":"确定要删除？","editorSetting":{"behavior":"delete"},"id":"u:cee85dad465d","label":"删除","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${TRIM(id)}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/menus/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"},{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}],"weight":0}},"type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_menu.delete\"]}"}],"id":"u:01259e435906","label":"操作","type":"operation"}],"draggable":true,"filter":{"actions":[{"actionType":"clear-and-submit","id":"u:cdb68b8d4b45","label":"重置","type":"button"},{"id":"u:253f267c0d95","label":"搜索","level":"primary","type":"submit"}],"body":[{"clearable":true,"id":"u:7fed8045c4d1","label":"名称","name":"name","size":"md","type":"input-text"},{"clearable":true,"id":"u:7fed8045c4d1","label":"路径","name":"path","size":"md","type":"input-text"}],"feat":"Insert","id":"u:ba2b6c5f28ea","panelClassName":"base-filter","title":"","type":"form"},"filterDefaultVisible":false,"filterTogglable":true,"footerToolbar":[{"type":"statistics"}],"headerToolbar":[{"actionType":"drawer","drawer":{"actionType":"drawer","actions":[{"actionType":"cancel","id":"u:1fcc9d820da8","label":"取消","type":"button"},{"actionType":"submit","id":"u:fbf3f20feca6","label":"提交","level":"primary","type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/menus"},"body":[{"clearable":true,"id":"u:1900aea073f5","label":"名称","name":"name","required":true,"row":0,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:d264b42b8dd9","label":"路径","name":"path","required":true,"row":1,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"description":"在 <a href=\"https://icones.js.org/\" target=\"_blank\">icones</a> 中寻找可用图标","id":"u:29e367eae357","label":"图标","name":"icon","row":2,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"enableNodePath":false,"id":"u:667de3ffb021","initiallyOpen":true,"label":"父级","labelField":"name","multiple":false,"name":"parent_id","row":3,"searchable":true,"showIcon":false,"source":"/system/menus/parent_options","type":"tree-select","value":0,"valueField":"id"},{"description":"数字越大越靠前","id":"u:7a65abe3e509","keyboard":true,"label":"排序","max":999999999,"min":0,"name":"sort","required":true,"row":4,"showSteps":true,"step":1,"type":"input-number","value":0},{"clearable":true,"id":"u:9341a4d17c70","label":"页面","multiple":false,"name":"page_sign","row":5,"searchable":true,"source":{"adaptor":"","messages":{},"method":"get","requestAdaptor":"","url":"/system/menus/page_options"},"type":"select"},{"description":"是否作为菜单显示","falseValue":0,"id":"u:171689bb598e","label":"是否显示","name":"visible","option":"","row":6,"trueValue":1,"type":"switch","value":1},{"animations":{},"body":[],"id":"u:8d075fa01421","isFixedHeight":false,"isFixedWidth":false,"isFreeContainer":true,"row":7,"size":"sm","style":{"position":"relative"},"type":"container","wrapperBody":false},{"body":[{"description":"登录后跳转的页面","falseValue":0,"id":"u:5c1b6af56b88","label":"首页","name":"is_home","option":"","row":7,"trueValue":1,"type":"switch","value":0},{"colSize":"1","description":"开启后, 该页面将不展示菜单栏","falseValue":0,"id":"u:54f5edb96e5c","label":"全屏","name":"is_full","option":"","row":7,"trueValue":1,"type":"switch"},{"colSize":"1","description":"开启后页面将缓存，重新打开时不会重新加载","falseValue":0,"id":"u:09899b7c6e71","label":"缓存页面","name":"keep_alive","option":"","row":7,"trueValue":1,"type":"switch"}],"collapsable":true,"collapsed":true,"id":"u:161a1a7915d1","row":8,"subFormMode":"","title":"更多设置","type":"fieldset"}],"dsType":"api","feat":"Insert","id":"u:f00289ce66a4","labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"},{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}]}},"resetAfterSubmit":true,"title":"新增数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"hideActions":false,"id":"u:5474e87e393d","resizable":false,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"新增","type":"drawer"},"editorSetting":{"behavior":"create"},"icon":"fa fa-plus","id":"u:2f6d4463cb00","label":"新增","level":"primary","type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_menu.create\"]}"},{"type":"bulk-actions"},{"align":"right","id":"u:df448b218def","type":"reload"},{"align":"right","id":"u:6046fb0268f2","tpl":"内容","type":"filter-toggler","wrapperComponent":""}],"id":"u:563375ef114c","itemActions":[],"messages":{},"onEvent":{"quickSaveItemSucc":{"actions":[{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}],"weight":0},"saveOrderSucc":{"actions":[{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}],"weight":0}},"perPage":20,"perPageAvailable":[5,10,20,50,100],"quickSaveItemApi":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/menus/quick_save"},"saveOrderApi":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/menus/sort"},"syncLocation":false,"type":"crud"}],"data":{"showFilter":false},"definitions":{},"id":"u:a8979bbc7044","initApi":{"method":"get","url":"/permissions"},"pullRefresh":{"disabled":true},"regions":["body"],"type":"page"}`),
		},
		{
			Name:   "权限",
			Sign:   "admin_permission",
			Schema: json.RawMessage(`{"body":[{"alwaysShowPagination":true,"api":{"method":"get","url":"/system/permissions"},"bulkActions":[{"actionType":"ajax","api":"","confirmText":"确定要删除？","editorSetting":{"behavior":"bulkDelete"},"icon":"fa fa-trash-o","id":"u:394fe8eeb9b7","label":"删除","level":"danger","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${ids}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/permissions/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"},{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}],"weight":0}},"type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_permission.bulk_delete\"]}"}],"columns":[{"id":"u:154e022f601d","label":"ID","name":"id","placeholder":"-","sortable":true,"type":"text"},{"className":"text-center h-full","id":"u:f6bc40b7d0db","items":[{"body":[{"className":"mr-3 text-xl h-full","icon":"${icon}","id":"u:0bdbe2fa7da7","type":"custom-svg-icon"}],"id":"u:36e21992b2fe","size":"none","type":"wrapper"},{"id":"u:fdfb43c65c06","tpl":"${name}","type":"tpl"}],"justify":"start","label":"名称","name":"name","placeholder":"-","type":"flex"},{"id":"u:acc2fc18bb5d","label":"标识","name":"sign","placeholder":"-","type":"text"},{"id":"u:0830ba2402e4","inline":true,"label":"排序","name":"sort","placeholder":"-","tpl":"","type":"tpl","wrapperComponent":""},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:05ee8b110a15","label":"创建时间","name":"created_at","placeholder":"-","sortable":true,"type":"date"},{"format":"YYYY-MM-DD HH:mm:ss","id":"u:a17c366510d7","label":"更新时间","name":"updated_at","placeholder":"-","sortable":true,"type":"date"},{"buttons":[{"actionType":"drawer","drawer":{"actionType":"drawer","actions":[{"actionType":"cancel","id":"u:9a6baf18c2c8","label":"取消","type":"button"},{"actionType":"confirm","id":"u:ce271dcff20b","label":"确定","primary":true,"type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/permissions"},"autoFocus":true,"body":[{"id":"u:b8e7e113f919","label":"ID","name":"id","row":-1,"type":"input-text","visible":false},{"clearable":true,"id":"u:1900aea073f5","label":"名称","name":"name","required":true,"row":0,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:d264b42b8dd9","label":"标识","name":"sign","required":true,"row":1,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"enableNodePath":false,"id":"u:667de3ffb021","initiallyOpen":true,"label":"父级","labelField":"name","multiple":false,"name":"parent_id","nodeBehavior":["check"],"row":3,"searchable":true,"showIcon":false,"source":"/system/permissions/parent_options","themeCss":{"actionControlClassName":{"marginLeft":""}},"type":"tree-select","value":0,"valueField":"id"},{"description":"数字越大越靠前","id":"u:7a65abe3e509","keyboard":true,"label":"排序","max":999999999,"min":0,"name":"sort","required":true,"row":4,"showSteps":true,"step":1,"type":"input-number","value":0},{"autoCheckChildren":false,"checkAll":true,"checkAllLabel":"全选","clearable":true,"defaultCheckAll":false,"enableNodePath":false,"id":"u:dafd5f0336f6","initiallyOpen":true,"joinValues":true,"label":"菜单","labelField":"name","multiple":true,"name":"menu_ids","nodeBehavior":["check"],"row":5,"searchable":true,"showIcon":false,"source":{"adaptor":"","messages":{},"method":"get","requestAdaptor":"","url":"/system/menus/parent_options"},"themeCss":{"actionControlClassName":{"marginLeft":""}},"type":"tree-select","valueField":"id","withChildren":true},{"addBtn":{"icon":"fa fa-plus","id":"u:aba832005742","label":"新增","level":"primary","size":"sm"},"addable":true,"description":"","draggable":true,"flat":false,"id":"u:1a0ebcfcc971","items":[{"id":"u:e94de4c5b633","name":"value","required":true,"type":"input-text","unique":true}],"joinValues":true,"label":"Api","labelRemark":{"className":"Remark--warning","content":"[请求方式] : [接口路径]\u003cbr\u003e\u003cbr\u003e示例:\u003cbr\u003e/system/users (任意请求方式)\u003cbr\u003eget:/system/permmission\u003cbr\u003epost:/system/menus\u003cbr\u003e\u003cbr\u003e使用正则匹配:\u003cbr\u003e^post:/system/menus","icon":"fa fa-question-circle","placement":"top","title":"规则说明","trigger":["hover"]},"multiple":true,"name":"api","removable":true,"removableMode":"icon","row":6,"strictMode":false,"syncFields":[],"type":"combo"}],"dsType":"api","feat":"Edit","id":"u:f00289ce66a4","initApi":{"adaptor":"","data":{"id":"${id}"},"messages":{},"method":"get","requestAdaptor":"","url":"/system/permissions/detail"},"labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"},{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}]}},"resetAfterSubmit":true,"title":"新增数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":false,"id":"u:cc540ca818a8","resizable":false,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"编辑","type":"drawer"},"editorSetting":{"behavior":"update"},"id":"u:29435e16d2df","label":"编辑","level":"link","type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_permission.update\"]}"},{"actionType":"","api":{"method":"post","url":"/system/permissions"},"className":"text-danger","confirmText":"确定要删除？","editorSetting":{"behavior":"delete"},"id":"u:cee85dad465d","label":"删除","level":"link","onEvent":{"click":{"actions":[{"actionType":"ajax","api":{"adaptor":"","data":{"ids":"${TRIM(id)}"},"messages":{},"method":"post","requestAdaptor":"","url":"/system/permissions/delete"},"ignoreError":false,"options":{},"outputVar":"responseResult"},{"actionType":"reload","componentId":"u:563375ef114c","groupType":"component"},{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}],"weight":0}},"type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_permission.delete\"]}"}],"id":"u:01259e435906","label":"操作","type":"operation"}],"draggable":true,"filter":{"actions":[{"actionType":"clear-and-submit","id":"u:cdb68b8d4b45","label":"重置","type":"button"},{"id":"u:253f267c0d95","label":"搜索","level":"primary","type":"submit"}],"body":[{"clearable":true,"id":"u:7fed8045c4d1","label":"名称","name":"name","size":"md","type":"input-text"},{"clearable":true,"id":"u:7fed8045c4d1","label":"标识","name":"sign","size":"md","type":"input-text"}],"feat":"Insert","id":"u:ba2b6c5f28ea","panelClassName":"base-filter","title":"","type":"form"},"filterDefaultVisible":false,"filterTogglable":true,"footerToolbar":[{"type":"statistics"}],"headerToolbar":[{"actionType":"drawer","drawer":{"actionType":"drawer","actions":[{"actionType":"cancel","id":"u:1fcc9d820da8","label":"取消","type":"button"},{"actionType":"submit","id":"u:fbf3f20feca6","label":"提交","level":"primary","type":"button"}],"body":[{"actions":[{"actionType":"cancel","label":"取消","type":"button"},{"actionType":"submit","label":"提交","level":"primary","type":"button"}],"api":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/permissions"},"body":[{"clearable":true,"id":"u:1900aea073f5","label":"名称","name":"name","required":true,"row":0,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"id":"u:d264b42b8dd9","label":"标识","name":"sign","required":true,"row":1,"type":"input-text","validations":{"maxLength":255}},{"clearable":true,"enableNodePath":false,"id":"u:667de3ffb021","initiallyOpen":true,"label":"父级","labelField":"name","multiple":false,"name":"parent_id","nodeBehavior":["check"],"row":3,"searchable":true,"showIcon":false,"source":"/system/permissions/parent_options","themeCss":{"actionControlClassName":{"marginLeft":""}},"type":"tree-select","value":0,"valueField":"id"},{"description":"数字越大越靠前","id":"u:7a65abe3e509","keyboard":true,"label":"排序","max":999999999,"min":0,"name":"sort","required":true,"row":4,"showSteps":true,"step":1,"type":"input-number","value":0},{"autoCheckChildren":false,"checkAll":true,"checkAllLabel":"全选","clearable":true,"defaultCheckAll":false,"enableNodePath":false,"id":"u:dafd5f0336f6","initiallyOpen":true,"joinValues":true,"label":"菜单","labelField":"name","multiple":true,"name":"menu_ids","nodeBehavior":["check"],"row":5,"searchable":true,"showIcon":false,"source":{"adaptor":"","messages":{},"method":"get","requestAdaptor":"","url":"/system/menus/parent_options"},"themeCss":{"actionControlClassName":{"marginLeft":""}},"type":"tree-select","valueField":"id","withChildren":true},{"addBtn":{"icon":"fa fa-plus","id":"u:aba832005742","label":"新增","level":"primary","size":"sm"},"addable":true,"draggable":true,"flat":false,"id":"u:1a0ebcfcc971","items":[{"id":"u:e94de4c5b633","name":"value","required":true,"type":"input-text","unique":true}],"joinValues":true,"label":"Api","labelRemark":{"className":"Remark--warning","content":"[请求方式] : [接口路径]\u003cbr\u003e\u003cbr\u003e示例:\u003cbr\u003e/system/users (任意请求方式)\u003cbr\u003eget:/system/permmission\u003cbr\u003epost:/system/menus\u003cbr\u003e\u003cbr\u003e使用正则匹配:\u003cbr\u003e^post:/system/menus","icon":"fa fa-question-circle","placement":"top","title":"规则说明","trigger":["hover"]},"multiple":true,"name":"api","removable":true,"removableMode":"icon","row":6,"strictMode":false,"syncFields":[],"type":"combo"}],"dsType":"api","feat":"Insert","id":"u:f00289ce66a4","labelAlign":"top","mode":"flex","onEvent":{"submitSucc":{"actions":[{"actionType":"search","componentId":"u:ebed804f316f","groupType":"component"},{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}]}},"resetAfterSubmit":true,"title":"新增数据","type":"form"}],"closeOnEsc":true,"closeOnOutside":true,"draggable":true,"hideActions":false,"id":"u:5474e87e393d","resizable":false,"showCloseButton":true,"showErrorMsg":true,"showLoading":true,"title":"新增","type":"drawer"},"editorSetting":{"behavior":"create"},"icon":"fa fa-plus","id":"u:2f6d4463cb00","label":"新增","level":"primary","type":"button","visibleOn":"${isSuperAdmin || _permissions[\"system.admin_permission.create\"]}"},{"type":"bulk-actions"},{"align":"right","id":"u:df448b218def","type":"reload"},{"align":"right","id":"u:6046fb0268f2","type":"filter-toggler","wrapperComponent":""}],"id":"u:563375ef114c","itemActions":[],"loadDataOnce":true,"matchFunc":"","messages":{},"onEvent":{"quickSaveItemSucc":{"actions":[{"actionType":"custom","ignoreError":false,"script":"window.$owl.refreshRoutes()"}],"weight":0}},"pageField":"","perPage":"","perPageAvailable":[5,10,20,50,100],"perPageField":"","quickSaveItemApi":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/permissions/quick_save"},"saveOrderApi":{"adaptor":"","messages":{},"method":"post","requestAdaptor":"","url":"/system/permissions/sort"},"syncLocation":false,"type":"crud"}],"data":{"showFilter":false},"definitions":{},"id":"u:a8979bbc7044","initApi":{"method":"get","url":"/permissions"},"pullRefresh":{"disabled":true},"regions":["body"],"type":"page"}`),
		},
		{
			Name:   "个人中心",
			Sign:   "user_center",
			Schema: json.RawMessage(`{"body":[{"api":"post:/user","body":[{"type":"fieldset","title":"基础信息","collapsable":false,"body":[{"label":"头像","name":"avatar","receiver":{"method":"post","url":"/upload","requestAdaptor":"","adaptor":"","messages":{}},"type":"input-image","id":"u:b8e333d75acb","accept":".jpeg, .jpg, .png, .gif","uploadType":"fileReceptor","proxy":true,"multiple":false,"hideUploadButton":true,"autoUpload":true,"crop":{"viewMode":1},"cropQuality":0.7,"limit":false},{"label":"姓名","name":"name","required":true,"type":"input-text","id":"u:fcf91118a9d3"}],"id":"u:d9344b16e803","subFormMode":"normal","className":"","bodyClassName":"p-5"},{"type":"container","body":[],"style":{"position":"relative","display":"flex","inset":"auto","flexWrap":"nowrap","flexDirection":"column","alignItems":"flex-start","height":"30px","overflowY":"visible"},"size":"none","wrapperBody":false,"id":"u:5d356d18f149","isFixedHeight":true,"isFixedWidth":false},{"type":"fieldset","title":"修改密码","collapsable":true,"body":[{"label":"密码","name":"password","type":"input-password","id":"u:a27d9df8f66a"},{"label":"确认密码","name":"confirm_password","type":"input-password","id":"u:c9a7d672f6a2","validations":{"equalsField":"password"},"validationErrors":{"equalsField":"两次输入密码不一致"}},{"label":"旧密码","name":"old_password","type":"input-password","id":"u:898828473c04","required":true,"validateOnChange":false,"visibleOn":"${password}"}],"id":"u:0051aedbe481","collapsed":true,"subFormMode":"normal","className":"","bodyClassName":"p-5","headingClassName":""}],"initApi":"/user","mode":"horizontal","panelClassName":"px-48 m:px-0","title":"","type":"form","id":"u:1097b79b6baa","actions":[{"type":"submit","label":"提交","primary":true,"id":"u:b0bdb64b2420"}],"feat":"Edit","dsType":"api","labelAlign":"left","onEvent":{"submitSucc":{"weight":0,"actions":[{"ignoreError":false,"actionType":"custom","script":"setTimeout(() => window.location.reload(), 1500)"}]}}}],"type":"page","id":"u:c2d28fe870b3","asideResizor":false,"pullRefresh":{"disabled":true}}`),
		},
	})
}

// 填充权限
func seedPermissions() {
	if !isNull(models.AdminPermission{}) {
		return
	}

	// 填充菜单权限
	var menus []models.AdminMenu
	db.Model(&models.AdminMenu{}).Find(&menus)

	insertPermissions := make([]models.AdminPermission, 0) // 权限
	withMenus := make([]map[string]interface{}, 0)         // 菜单和权限关联
	for _, menu := range menus {
		insertPermissions = append(insertPermissions, models.AdminPermission{
			Name:     menu.Name,
			ParentId: menu.ParentId,
			Sign:     strings.ReplaceAll(strings.TrimLeft(menu.Path, "/"), "/", "."),
			Api:      `["get:` + strings.ReplaceAll(menu.Path, "admin_", "") + `s"]`,
		})

		withMenus = append(withMenus, map[string]interface{}{
			"admin_menu_id":       menu.ID,
			"admin_permission_id": menu.ID,
		})
	}

	db.Model(&models.AdminPermission{}).Create(insertPermissions)

	db.Exec("truncate table admin_menu_permission") // 清除关联表数据
	db.Table("admin_menu_permission").Create(withMenus)

	// 操作权限
	var permissions []models.AdminPermission
	db.Model(&models.AdminPermission{}).Find(&permissions)

	insertPermissions = make([]models.AdminPermission, 0)
	for _, permission := range permissions {
		if !strings.HasPrefix(permission.Sign, "system.") {
			continue
		}

		formApi := func() string {
			api := strings.ReplaceAll(permission.Api, "get:", "post:")

			if strings.HasSuffix(permission.Sign, "admin_user") {
				api = strings.ReplaceAll(api, `"]`, `", "get:/system/users/role_options"]`)
				api = strings.ReplaceAll(api, `"]`, `", "post:/system/users/quick_save"]`)
			}

			if strings.HasSuffix(permission.Sign, "admin_permission") {
				api = strings.ReplaceAll(api, `"]`, `", "get:/system/permissions/parent_options"]`)
				api = strings.ReplaceAll(api, `"]`, `", "post:/system/permissions/sort"]`)
			}

			if strings.HasSuffix(permission.Sign, "admin_menu") {
				api = strings.ReplaceAll(api, `"]`, `", "get:/system/menus/parent_options"]`)
				api = strings.ReplaceAll(api, `"]`, `", "get:/system/menus/page_options"]`)
				api = strings.ReplaceAll(api, `"]`, `", "post:/system/menus/quick_save"]`)
				api = strings.ReplaceAll(api, `"]`, `", "post:/system/menus/sort"]`)
			}

			if strings.HasSuffix(permission.Sign, "admin_permission") || strings.HasSuffix(permission.Sign, "admin_menu") {
				api = strings.ReplaceAll(api, `"]`, `", "get:/system/menus/parent_options"]`)
			}

			return api
		}()

		insertPermissions = append(insertPermissions, []models.AdminPermission{
			// 新增
			{
				Name:     "新增数据",
				ParentId: permission.ID,
				Sign:     permission.Sign + ".create",
				Api:      formApi,
			},
			// 修改
			{
				Name:     "获取详情",
				ParentId: permission.ID,
				Sign:     permission.Sign + ".show",
				Api:      strings.ReplaceAll(permission.Api, `"]`, `/detail"]`),
			},
			{
				Name:     "修改数据",
				ParentId: permission.ID,
				Sign:     permission.Sign + ".update",
				Api:      formApi,
			},
			// 删除
			{
				Name:     "删除数据",
				ParentId: permission.ID,
				Sign:     permission.Sign + ".delete",
				Api:      strings.ReplaceAll(strings.ReplaceAll(permission.Api, `"]`, `/delete"]`), "get:", "post:"),
			},
			// 批量删除
			{
				Name:     "批量删除数据",
				ParentId: permission.ID,
				Sign:     permission.Sign + ".bulk_delete",
				Api:      strings.ReplaceAll(strings.ReplaceAll(permission.Api, `"]`, `/delete"]`), "get:", "post:"),
			},
		}...)

		if strings.HasSuffix(permission.Sign, "admin_role") {
			insertPermissions = append(insertPermissions, models.AdminPermission{
				Name:     "设置权限",
				ParentId: permission.ID,
				Sign:     permission.Sign + ".edit_permission",
				Api: func() string {
					api := strings.ReplaceAll(permission.Api, `"]`, `", "get:/system/roles/permissions"]`)
					api = strings.ReplaceAll(api, `"]`, `", "post:/system/roles/permissions"]`)
					api = strings.ReplaceAll(api, `"]`, `", "get:/system/permissions/parent_options"]`)

					return api
				}(),
			})
		}
	}

	db.Model(&models.AdminPermission{}).Create(insertPermissions)
}
