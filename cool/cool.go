package cool

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/guid"
	"gorm.io/gorm"
)

var (
	GormDBS      map[string]*gorm.DB // 定义全局gorm.DB对象集合 仅供内部使用
	CacheEPS     = gcache.New()      // 定义全局缓存对象	供EPS使用
	CacheManager = gcache.New()      // 定义全局缓存对象	供其他业务使用
	ProcessFlag  = guid.S()          // 定义全局进程标识
)

type MgormDBS map[string]*gorm.DB

func init() {
	var (
		ctx         g.Ctx
		redisConfig = &gredis.Config{}
	)
	GormDBS = make(MgormDBS)
	g.Log().Debug(ctx, "cool init,初始化核心模块,请等待...")
	g.Log().Debug(ctx, "初始化缓存")
	redisVar, err := g.Cfg().Get(ctx, "redis.default")
	if err != nil {
		panic(err)
	}
	if !redisVar.IsEmpty() {
		redisVar.Struct(redisConfig)
		redis, err := gredis.New(redisConfig)
		if err != nil {
			panic(err)
		}
		CacheManager.SetAdapter(gcache.NewAdapterRedis(redis))
	}
	g.Log().Debug(ctx, "初始化缓存完成")
	g.Log().Debug(ctx, "当前实例ID:", ProcessFlag)
}

// cool.OK 正常返回
type BaseRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 返回正常结果
func Ok(data interface{}) *BaseRes {

	return &BaseRes{
		Code:    1000,
		Message: "success",
		Data:    data,
	}
}

// 失败返回结果
func Fail(message string) *BaseRes {
	return &BaseRes{
		Code:    1001,
		Message: message,
	}
}
