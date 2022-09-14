package cool

import "github.com/cool-team-official/cool-admin-go/cool/coolconfig"

var (
	Config            = coolconfig.Config            // 配置中的cool节相关配置
	GetCfgWithDefault = coolconfig.GetCfgWithDefault // GetCfgWithDefault 获取配置，如果配置不存在，则使用默认值
)
