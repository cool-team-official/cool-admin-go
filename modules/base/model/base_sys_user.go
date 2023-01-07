package model

import "github.com/cool-team-official/cool-admin-go/cool"

const TableNameBaseSysUser = "base_sys_user"

// BaseSysUser mapped from table <base_sys_user>
type BaseSysUser struct {
	*cool.Model
	DepartmentID uint    `gorm:"column:departmentId;type:bigint;index" json:"departmentId"`        // 部门ID
	Name         *string `gorm:"column:name;type:varchar(255)" json:"name"`                        // 姓名
	Username     string  `gorm:"column:username;type:varchar(100);not null;Index" json:"username"` // 用户名
	Password     string  `gorm:"column:password;type:varchar(255);not null" json:"password"`       // 密码
	PasswordV    *int32  `gorm:"column:passwordV;type:int;not null;default:1" json:"passwordV"`    // 密码版本, 作用是改完密码，让原来的token失效
	NickName     *string `gorm:"column:nickName;type:varchar(255)" json:"nickName"`                // 昵称
	HeadImg      *string `gorm:"column:headImg;type:varchar(255)" json:"headImg"`                  // 头像
	Phone        *string `gorm:"column:phone;type:varchar(20);index" json:"phone"`                 // 手机
	Email        *string `gorm:"column:email;type:varchar(255)" json:"email"`                      // 邮箱
	Status       *int32  `gorm:"column:status;not null;default:1" json:"status"`                   // 状态 0:禁用 1：启用
	Remark       *string `gorm:"column:remark;type:varchar(255)" json:"remark"`                    // 备注
	SocketID     *string `gorm:"column:socketId;type:varchar(255)" json:"socketId"`                // socketId
}

// TableName BaseSysUser's table name
func (*BaseSysUser) TableName() string {
	return TableNameBaseSysUser
}

// NewBaseSysUser 创建一个新的BaseSysUser
func NewBaseSysUser() *BaseSysUser {
	return &BaseSysUser{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&BaseSysUser{})
}
