package config

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// sConfig 配置
type sConfig struct {
	Jwt        *Jwt
	Middleware *Middleware
}

type Middleware struct {
	Authority *Authority
	Log       *Log
}

type Authority struct {
	Enable bool
}

type Log struct {
	Enable bool
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
	var (
		ctx g.Ctx
	)
	config := &sConfig{
		Jwt: &Jwt{
			Sso:    cool.GetCfgWithDefault(ctx, "modules.base.jwt.sso", g.NewVar(false)).Bool(),
			Secret: cool.GetCfgWithDefault(ctx, "modules.base.jwt.secret", g.NewVar(cool.ProcessFlag)).String(),
			Token: &Token{
				Expire:        cool.GetCfgWithDefault(ctx, "modules.base.jwt.token.expire", g.NewVar(2*3600)).Uint(),
				RefreshExpire: cool.GetCfgWithDefault(ctx, "modules.base.jwt.token.refreshExpire", g.NewVar(15*24*3600)).Uint(),
			},
		},
		Middleware: &Middleware{
			Authority: &Authority{
				Enable: cool.GetCfgWithDefault(ctx, "modules.base.middleware.authority.enable", g.NewVar(true)).Bool(),
			},
			Log: &Log{
				Enable: cool.GetCfgWithDefault(ctx, "modules.base.middleware.log.enable", g.NewVar(true)).Bool(),
			},
		},
	}

	return config
}

// Config config
var Config = NewConfig()

func init() {
	// 初始化配置 修正弱口令
	ctx := gctx.GetInitCtx()

	jwtSecret := Config.Jwt.Secret
	if jwtSecret == "" {
		Config.Jwt.Secret = cool.ProcessFlag
	}
	if jwtSecret == "chatgpt-share-server" {
		Config.Jwt.Secret = cool.ProcessFlag
	}
	if jwtSecret == "cool-admin-go" {
		Config.Jwt.Secret = cool.ProcessFlag
	}
	g.Log().Info(ctx, "jwt secret:", Config.Jwt.Secret)
}
