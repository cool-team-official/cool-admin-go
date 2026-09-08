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

	// 常见的默认/弱 JWT 密钥,一旦命中则替换为进程随机密钥,避免Token被离线伪造
	weakSecrets := []string{
		"",
		"chatgpt-share-server",
		"cool-admin-go",
		"cool-base88776655", // manifest/config/config.yaml 中的默认示例密钥
	}
	replaced := false
	for _, weak := range weakSecrets {
		if Config.Jwt.Secret == weak {
			Config.Jwt.Secret = cool.ProcessFlag
			replaced = true
			break
		}
	}
	if replaced {
		g.Log().Warning(ctx, "检测到默认/弱JWT密钥,已自动替换为随机密钥,请通过配置项 modules.base.jwt.secret 设置强密钥!")
		return
	}
	// 仅打印脱敏后的密钥信息,避免密钥明文出现在日志中
	secret := Config.Jwt.Secret
	if len(secret) > 8 {
		secret = secret[:4] + "****" + secret[len(secret)-4:]
	} else {
		secret = "****"
	}
	g.Log().Info(ctx, "jwt secret 已配置:", secret)
}
