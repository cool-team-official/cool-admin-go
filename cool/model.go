package cool

import (
	"time"

	"gorm.io/gorm"
)

type IModel interface {
	TableName() string
	GroupName() string
}
type Model struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreateTime time.Time      `gorm:"index" json:"createTime"`
	UpdateTime time.Time      `gorm:"index" json:"UpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

// 返回表名
func (m *Model) TableName() string {
	return "this_table_should_not_exist"
}

// 返回分组名
func (m *Model) GroupName() string {
	return "default"
}

func NewModel() *Model {
	return &Model{
		ID:         0,
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
		DeletedAt:  gorm.DeletedAt{},
	}
}
