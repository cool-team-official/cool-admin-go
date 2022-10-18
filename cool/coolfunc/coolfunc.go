package coolfunc

import "github.com/gogf/gf/v2/frame/g"

type CoolFunc interface {
	// Func handler
	Func(ctx g.Ctx, param string) error
	// IsSingleton 是否单例,当为true时，只能有一个任务在执行,在注意函数为计划任务时使用
	IsSingleton() bool
	// IsAllWorker 是否所有worker都执行
	IsAllWorker() bool
}
