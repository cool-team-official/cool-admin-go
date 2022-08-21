package cool

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type Admin struct {
	IsRefresh       bool   `json:"isRefresh"`
	RoleIds         []int  `json:"roleIds"`
	Username        string `json:"username"`
	UserId          uint   `json:"userId"`
	PasswordVersion *int32 `json:"passwordVersion"`
}

// 获取传入ctx 中的 admin 对象
func GetAdmin(ctx context.Context) *Admin {
	r := g.RequestFromCtx(ctx)
	admin := &Admin{}
	err := gjson.New(r.GetCtxVar("admin").String()).Scan(admin)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return admin
}
