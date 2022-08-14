package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/util/gmeta"
)

type SampleAddReq struct {
	cool.AddReq
}
type Sample struct {
	gmeta.Meta `group:"default" tableName:"xxxx_sample"`
	cool.Model
	SampleAddReq
}

// 定义表名
func (Sample) TableName() string {
	return gmeta.Get(Sample{}, "tableName").String()
}

// NewSample 创建一个新的Sample实例
func NewSample() *Sample {
	return &Sample{}
}

// 初始化
func init() {
	cool.CreateTable(&Sample{})
}
