package cool

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"gorm.io/gorm"
)

var (
	GromDBS MGromDBS // 定义全局gorm.DB对象集合
	// CacheEPS *gcache.Cache // 定义全局缓存对象	供EPS使用
	CacheEPS = gcache.New() // 定义全局缓存对象	供EPS使用
	Cache    = gcache.New() // 定义全局缓存对象	供其他业务使用
)

type MGromDBS map[string]*gorm.DB

func init() {
	var (
		ctx         g.Ctx
		redisConfig = &gredis.Config{}
	)
	GromDBS = make(MGromDBS)
	g.Log().Debug(ctx, "cool init,初始化核心模块,请等待...")
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
		Cache.SetAdapter(gcache.NewAdapterRedis(redis))
	}
}

// cool.OK 正常返回
type BaseRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Ok(data interface{}) *BaseRes {

	return &BaseRes{
		Code:    1000,
		Message: "success",
		Data:    data,
	}
}

// GDBModel 数据库连接
func GDBModel(m IModel) *gdb.Model {
	return g.DB(m.GroupName()).Model(m.TableName()).Clone()
}
