package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
)

const TableNameDictType = "dict_type"

// DictType mapped from table <dict_type>
type DictType struct {
	*cool.Model
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"` // 名称
	Key  string `gorm:"column:key;type:varchar(255);not null" json:"key"`   // 标识
}

// TableName DictType's table name
func (*DictType) TableName() string {
	return TableNameDictType
}

// GroupName DictType's table group
func (*DictType) GroupName() string {
	return "default"
}

// NewDictType create a new DictType
func NewDictType() *DictType {
	return &DictType{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&DictType{})
}
