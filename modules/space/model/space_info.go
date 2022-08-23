package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
)

const TableNameSpaceInfo = "space_info"

// SpaceInfo mapped from table <space_info>
type SpaceInfo struct {
	*cool.Model
	URL        string `gorm:"column:url;type:varchar(255);not null;comment:地址" json:"url"`   // 地址
	Type       string `gorm:"column:type;type:varchar(255);not null;comment:类型" json:"type"` // 类型
	ClassifyID *int64 `gorm:"column:classifyId;type:bigint;comment:分类ID" json:"classifyId"`  // 分类ID
}

// TableName SpaceInfo's table name
func (*SpaceInfo) TableName() string {
	return TableNameSpaceInfo
}

// GroupName SpaceInfo's table group
func (*SpaceInfo) GroupName() string {
	return "default"
}

// NewSpaceInfo create a new SpaceInfo
func NewSpaceInfo() *SpaceInfo {
	return &SpaceInfo{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&SpaceInfo{})
}
