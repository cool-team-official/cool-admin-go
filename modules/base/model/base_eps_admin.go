package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/util/gmeta"
)

type Eps struct {
	gmeta.Meta `group:"default" tableName:"base_eps_admin"`
	Id         int    `json:"id"`
	Module     string `json:"module" field:"module"`
	Method     string // 请求方法 例如：GET
	Path       string // 请求路径 例如：/welcome
	Prefix     string // 路由前缀 例如：/admin/base/open
	Summary    string // 描述 例如：欢迎页面
	Tag        string // 标签 例如：base  好像暂时不用
	Dts        string // 未知 例如：{} 好像暂时不用
}

// 定义表名
func (Eps) TableName() string {
	return gmeta.Get(Eps{}, "tableName").String()
}

// 定义分组
func (Eps) GroupName() string {
	return "default"
}

// NewEps 创建一个新的Eps实例
func NewEps() *Eps {
	return &Eps{}
}

// 初始化
func init() {
	cool.CreateTable(&Eps{})
}
