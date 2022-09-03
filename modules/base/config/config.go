package config

import "github.com/gogf/gf/v2/frame/g"

// sConfig 配置
type sConfig struct {
	Jwt struct {
		Sso    bool   `json:"sso"`
		Secret string `json:"secret"`
		Token  struct {
			Expire        uint `json:"expire"`
			RefreshExpire uint `json:"refreshExprire"`
		} `json:"token"`
	} `json:"jwt"`
}

// NewConfig new config
func NewConfig() *sConfig {
	var ctx g.Ctx
	config := &sConfig{}
	g.Cfg().MustGet(ctx, "modules.base").Struct(config)
	return config
}

// Config config
var Config = NewConfig()
