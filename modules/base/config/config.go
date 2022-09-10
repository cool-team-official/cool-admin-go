package config

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/frame/g"
)

// sConfig 配置
type sConfig struct {
	Jwt *Jwt
}

type Token struct {
	Expire        uint `json:"expire"`
	RefreshExpire uint `json:"refreshExprire"`
}

type Jwt struct {
	Sso    bool   `json:"sso"`
	Secret string `json:"secret"`
	Token  *Token `json:"token"`
}

// NewConfig new config
func NewConfig() *sConfig {
	var ctx g.Ctx
	config := &sConfig{
		Jwt: &Jwt{
			Sso:    cool.GetCfgWithDefault(ctx, "modules.base.jwt.sso", g.NewVar(false)).Bool(),
			Secret: cool.GetCfgWithDefault(ctx, "modules.base.jwt.secret", g.NewVar("cool-admin-go")).String(),
			Token: &Token{
				Expire:        cool.GetCfgWithDefault(ctx, "modules.base.jwt.token.expire", g.NewVar(2*3600)).Uint(),
				RefreshExpire: cool.GetCfgWithDefault(ctx, "modules.base.jwt.token.refreshExpire", g.NewVar(15*24*3600)).Uint(),
			},
		},
	}
	return config
}

// Config config
var Config = NewConfig()
