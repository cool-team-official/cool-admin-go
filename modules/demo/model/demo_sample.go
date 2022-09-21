package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
)

const TableNameDemoSample = "demo_sample"

// DemoSample mapped from table <demo_sample>
type DemoSample struct {
	*cool.Model
	// Name string `gorm:"column:name;not null;comment:名称" json:"name"`
}

// TableName DemoSample's table name
func (*DemoSample) TableName() string {
	return TableNameDemoSample
}

// GroupName DemoSample's table group
func (*DemoSample) GroupName() string {
	return "default"
}

// NewDemoSample create a new DemoSample
func NewDemoSample() *DemoSample {
	return &DemoSample{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&DemoSample{})
}
