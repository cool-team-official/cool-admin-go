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
	CreateTime *time.Time     `gorm:"column:createTime;not null;index,priority:1;autoCreateTime:nano;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime *time.Time     `gorm:"column:updateTime;not null;index,priority:1;autoUpdateTime:nano;comment:更新时间" json:"updateTime"` // 更新时间
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
		CreateTime: &time.Time{},
		UpdateTime: &time.Time{},
		DeletedAt:  gorm.DeletedAt{},
	}
}
