# 页面Schema管理系统

本目录包含程序内定义的页面Schema文件，这些文件会在编译时嵌入到应用程序中。

## 功能特性

- **程序内定义**: 将页面Schema定义为JSON文件，使用Go的`embed`包嵌入到程序中
- **优先级机制**: 程序内定义的Schema优先于数据库存储的Schema
- **自动加载**: 应用启动时自动加载所有JSON文件
- **向后兼容**: 完全兼容现有的数据库存储方式

## 目录结构

```
internal/schema/
├── README.md           # 本说明文档
├── manager.go          # Schema管理器
├── dashboard.json      # 控制台Schema
├── login.json          # 登录页面Schema
├── admin_page.json     # 页面管理Schema
├── admin_user.json     # 管理员Schema
├── admin_role.json     # 角色Schema
├── admin_menu.json     # 菜单Schema
├── admin_permission.json # 权限Schema
└── [其他页面].json     # 其他页面Schema文件
```

## 使用方法

### 1. 添加新的页面Schema

在此目录下创建新的JSON文件，文件名即为页面的`sign`标识：

```json
{
  "type": "page",
  "title": "页面标题",
  "body": {
    // 页面内容定义
  }
}
```

### 2. Schema文件命名规则

- 文件名格式：`{page_sign}.json`
- 例如：`user_list.json` 对应页面标识 `user_list`

### 3. 程序中使用

```go
import "amis-base/internal/schema"

// 获取指定页面的Schema
schemaStr := schema.GetSchemaBySign("admin_page")

// 检查Schema是否存在
exists := schema.HasSchema("admin_page")

// 获取所有程序内定义的Schema信息
allSchemas := schema.GetAllSchemas()
```

## 优先级机制

当请求页面Schema时，系统按以下优先级查找：

1. **程序内定义** - 首先检查`internal/schema/`目录中的JSON文件
2. **数据库存储** - 如果程序内不存在，则查询数据库中的`admin_pages`表

这种机制确保：
- 核心页面可以通过程序内定义保证稳定性
- 用户自定义页面仍可通过数据库灵活管理
- 程序内定义的页面不会被意外修改或删除

## API变更

### PageSchema接口

`GET /admin/page_schema?sign={page_sign}`

现在会优先返回程序内定义的Schema，如果不存在则回退到数据库查找。

### PageOptions接口

`GET /admin/system/menus/page_options`

返回的页面选项列表现在包含：
1. 所有程序内定义的页面
2. 数据库中不与程序内定义重复的页面

## 注意事项

1. **文件格式**: 所有Schema文件必须是有效的JSON格式
2. **编译时嵌入**: 修改JSON文件后需要重新编译应用程序
3. **唯一性**: 程序内定义的页面标识不应与数据库中的重复
4. **性能**: 程序内定义的Schema在内存中缓存，访问速度更快

## 示例Schema

参考目录中的各个JSON文件了解Schema的结构和格式。

## 迁移完成

已成功将以下页面从 `internal/pkg/db/seed.go` 迁移到程序内定义：

- **dashboard** - 控制台页面
- **login** - 登录页面
- **admin_page** - 页面管理
- **admin_user** - 管理员管理
- **admin_role** - 角色管理
- **admin_menu** - 菜单管理
- **admin_permission** - 权限管理

这些页面现在优先使用程序内定义的Schema，确保了核心功能的稳定性和一致性。
