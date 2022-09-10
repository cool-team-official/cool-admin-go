package cool

import "github.com/gogf/gf/v2/frame/g"

type sConfig struct {
	AutoMigrate bool
	File        *file `json:"file"`
}
type file struct {
	Mode   string `json:"mode"`
	Domain string `json:"domain"`
}

// NewConfig new config
func newConfig() *sConfig {
	var ctx g.Ctx
	config := &sConfig{
		AutoMigrate: GetCfgWithDefault(ctx, "cool.autoMigrate", g.NewVar(false)).Bool(),
		File: &file{
			Mode:   GetCfgWithDefault(ctx, "cool.file.mode", g.NewVar("local")).String(),
			Domain: GetCfgWithDefault(ctx, "cool.file.domain", g.NewVar("http://127.0.0.1:8300")).String(),
		},
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
	if value.IsEmpty() {
		return defaultValue
	}
	return value
}
