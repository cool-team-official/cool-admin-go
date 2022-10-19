package coolconfig

import "github.com/gogf/gf/v2/frame/g"

// cool config
type sConfig struct {
	AutoMigrate bool  `json:"auto_migrate,omitempty"` // 是否自动创建表
	Eps         bool  `json:"eps,omitempty"`          // 是否开启eps
	File        *file `json:"file,omitempty"`         // 文件上传配置
}

// 文件上传配置
type file struct {
	Mode   string `json:"mode"`   // 模式 local oss
	Domain string `json:"domain"` // 域名 http://
}

// NewConfig new config
func newConfig() *sConfig {
	var ctx g.Ctx
	config := &sConfig{
		AutoMigrate: GetCfgWithDefault(ctx, "cool.autoMigrate", g.NewVar(false)).Bool(),
		Eps:         GetCfgWithDefault(ctx, "cool.eps", g.NewVar(false)).Bool(),
		File: &file{
			Mode:   GetCfgWithDefault(ctx, "cool.file.mode", g.NewVar("none")).String(),
			Domain: GetCfgWithDefault(ctx, "cool.file.domain", g.NewVar("http://127.0.0.1:8300")).String()},
	}
	return config
}

// Config config
var Config = newConfig()

// GetCfgWithDefault get config with default value
func GetCfgWithDefault(ctx g.Ctx, key string, defaultValue *g.Var) *g.Var {
	value, err := g.Cfg().Get(ctx, key)
	if err != nil {
		return defaultValue
	}
	if value.IsEmpty() || value.IsNil() {
		return defaultValue
	}
	return value
}
