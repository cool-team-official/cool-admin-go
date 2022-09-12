package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
)

const TableNameDemoGoods = "demo_goods"

// DemoGoods mapped from table <demo_goods>
type DemoGoods struct {
	*cool.Model
	Name string `gorm:"not null" json:"name"`
}

// TableName DemoGoods's table name
func (*DemoGoods) TableName() string {
	return TableNameDemoGoods
}

// GroupName DemoGoods's table group
func (*DemoGoods) GroupName() string {
	return "default"
}

// NewDemoGoods create a new DemoGoods
func NewDemoGoods() *DemoGoods {
	return &DemoGoods{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&DemoGoods{})
}
