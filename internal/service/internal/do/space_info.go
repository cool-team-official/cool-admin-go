// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaceInfo is the golang structure of table space_info for DAO operations like Where/Data.
type SpaceInfo struct {
	g.Meta     `orm:"table:space_info, do:true"`
	Id         interface{} // ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	Url        interface{} // 地址
	Type       interface{} // 类型
	ClassifyId interface{} // 分类ID
}
