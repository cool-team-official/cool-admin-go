package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
)

const TableNameBaseEpsApp = "base_eps_app"

// BaseEpsApp mapped from table <base_eps_app>
type BaseEpsApp struct {
	Id      int    `json:"id"`
	Module  string `json:"module" field:"module"`
	Method  string // 请求方法 例如：GET
	Path    string // 请求路径 例如：/welcome
	Prefix  string // 路由前缀 例如：/admin/base/open
	Summary string // 描述 例如：欢迎页面
	Tag     string // 标签 例如：base  好像暂时不用
	Dts     string // 未知 例如：{} 好像暂时不用
}

// TableName BaseEpsApp's table name
func (*BaseEpsApp) TableName() string {
	return TableNameBaseEpsApp
}

// GroupName BaseEpsApp's table group
func (*BaseEpsApp) GroupName() string {
	return "default"
}

// NewBaseEpsApp create a new BaseEpsApp
func NewBaseEpsApp() *BaseEpsApp {
	return &BaseEpsApp{}
}

// init 创建表
func init() {
	cool.CreateTable(&BaseEpsApp{})
}
