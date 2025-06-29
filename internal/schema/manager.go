package schema

import (
	"embed"
	"encoding/json"
	"strings"
)

//go:embed *.json
var schemaFiles embed.FS

// SchemaInfo 程序内定义的schema信息
type SchemaInfo struct {
	Sign   string          `json:"sign"`
	Name   string          `json:"name"`
	Schema json.RawMessage `json:"schema"`
}

// Manager schema管理器
type Manager struct {
	schemas map[string]*SchemaInfo
}

// NewManager 创建新的schema管理器
func NewManager() *Manager {
	m := &Manager{
		schemas: make(map[string]*SchemaInfo),
	}
	m.loadEmbeddedSchemas()
	return m
}

// loadEmbeddedSchemas 加载嵌入的schema文件
func (m *Manager) loadEmbeddedSchemas() {
	// 需要手动开启
	return

	entries, err := schemaFiles.ReadDir(".")
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		// 从文件名提取sign（去掉.json后缀）
		sign := strings.TrimSuffix(entry.Name(), ".json")

		// 读取文件内容
		content, err := schemaFiles.ReadFile(entry.Name())
		if err != nil {
			continue
		}

		// 解析JSON以获取title作为name
		var schemaData map[string]any
		if err := json.Unmarshal(content, &schemaData); err != nil {
			continue
		}

		name := sign // 默认使用sign作为name
		if title, ok := schemaData["title"].(string); ok && title != "" {
			name = title
		}

		// 注册schema
		m.schemas[sign] = &SchemaInfo{
			Sign:   sign,
			Name:   name,
			Schema: json.RawMessage(content),
		}
	}
}

// GetSchemaBySign 根据sign获取schema
func (m *Manager) GetSchemaBySign(sign string) string {
	if schema, exists := m.schemas[sign]; exists {
		return string(schema.Schema)
	}
	return ""
}

// HasSchema 检查是否存在指定sign的schema
func (m *Manager) HasSchema(sign string) bool {
	_, exists := m.schemas[sign]
	return exists
}

// GetAllSchemas 获取所有程序内定义的schema信息
func (m *Manager) GetAllSchemas() []*SchemaInfo {
	schemas := make([]*SchemaInfo, 0, len(m.schemas))
	for _, schema := range m.schemas {
		schemas = append(schemas, schema)
	}
	return schemas
}

// GetSchemaInfo 获取schema信息
func (m *Manager) GetSchemaInfo(sign string) *SchemaInfo {
	return m.schemas[sign]
}

// 全局schema管理器实例
var globalManager *Manager

// GetManager 获取全局schema管理器
func GetManager() *Manager {
	if globalManager == nil {
		globalManager = NewManager()
	}
	return globalManager
}

// GetSchemaBySign 全局函数：根据sign获取schema
func GetSchemaBySign(sign string) string {
	return GetManager().GetSchemaBySign(sign)
}

// HasSchema 全局函数：检查是否存在指定sign的schema
func HasSchema(sign string) bool {
	return GetManager().HasSchema(sign)
}

// GetAllSchemas 全局函数：获取所有程序内定义的schema信息
func GetAllSchemas() []*SchemaInfo {
	return GetManager().GetAllSchemas()
}
