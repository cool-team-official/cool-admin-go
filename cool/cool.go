package cool

import (
	"gorm.io/gorm"
)

// 定义全局gorm.DB对象集合
var GromDBS MGromDBS

type MGromDBS map[string]*gorm.DB

func init() {
	GromDBS = make(MGromDBS)
}
