package cool

import "github.com/gogf/gf/v2/frame/g"

type sConfig struct {
	File *file `json:"file"`
}
type file struct {
	Mode   string `json:"mode"`
	Domain string `json:"domain"`
}

// NewConfig new config
func newConfig() *sConfig {
	var ctx g.Ctx
	config := &sConfig{
		File: &file{
			Mode:   "local",
			Domain: "http://127.0.0.1",
		},
	}
	c, _ := g.Cfg().Get(ctx, "cool")
	if c != nil {
		c.Struct(config)
	}
	return config
}

// Config config
var Config = newConfig()
