package model

import "github.com/cool-team-official/cool-admin-go/cool"

const TableNameBaseSysParam = "base_sys_param"

// BaseSysParam mapped from table <base_sys_param>
type BaseSysParam struct {
	*cool.Model
	KeyName  string  `gorm:"column:keyName;type:varchar(255);not null;index:IDX_cf19b5e52d8c71caa9c4534454,priority:1" json:"keyName"` // 键位
	Name     string  `gorm:"column:name;type:varchar(255);not null" json:"name"`                                                       // 名称
	Data     string  `gorm:"column:data;type:text;not null" json:"data"`                                                               // 数据
	DataType int32   `gorm:"column:dataType;not null;default:0" json:"dataType"`                                                       // 数据类型 0:字符串 1：数组 2：键值对
	Remark   *string `gorm:"column:remark;type:varchar(255)" json:"remark"`                                                            // 备注
}

// TableName BaseSysParam's table name
func (*BaseSysParam) TableName() string {
	return TableNameBaseSysParam
}

// NewBaseSysParam 创建一个新的BaseSysParam
func NewBaseSysParam() *BaseSysParam {
	return &BaseSysParam{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&BaseSysParam{})
}
