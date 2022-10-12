package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type BaseSysLogService struct {
	*cool.Service
}

func NewBaseSysLogService() *BaseSysLogService {
	return &BaseSysLogService{
		&cool.Service{
			Model: model.NewBaseSysLog(),
			PageQueryOp: &cool.QueryOp{
				FieldEQ:      []string{},
				KeyWordField: []string{},
				AddOrderby:   map[string]string{},

				Select: "base_sys_log.*,user.name ",
				Join: []*cool.JoinOp{
					{
						Model:     model.NewBaseSysUser(),
						Alias:     "user",
						Type:      "LeftJoin",
						Condition: "user.id = base_sys_log.userID",
					},
				},
			},
		},
	}
}

// Record 记录日志
func (s *BaseSysLogService) Record(ctx g.Ctx) {
	var (
		admin = cool.GetAdmin(ctx)
		r     = g.RequestFromCtx(ctx)
	)
	baseSysLog := model.NewBaseSysLog()
	baseSysLog.UserID = admin.UserId
	baseSysLog.Action = r.Method + ":" + r.URL.Path
	baseSysLog.IP = r.GetClientIp()
	baseSysLog.IPAddr = r.GetClientIp()
	baseSysLog.Params = r.GetBodyString()
	// g.DumpWithType(baseSysLog)
	m := cool.DBM(s.Model)
	m.Insert(g.Map{
		"userId": baseSysLog.UserID,
		"action": baseSysLog.Action,
		"ip":     baseSysLog.IP,
		"ipAddr": baseSysLog.IPAddr,
		"params": baseSysLog.Params,
	})
}

// Clear 清除日志
func (s *BaseSysLogService) Clear(isAll bool) (err error) {
	BaseSysConfService := NewBaseSysConfService()
	m := cool.DBM(s.Model)
	if isAll {
		_, err = m.Delete("1=1")
	} else {
		keepDays := gconv.Int(BaseSysConfService.GetValue("logKeep"))
		_, err = m.Delete("create_time < ?", gtime.Now().AddDate(0, 0, -keepDays).String())
	}
	return
}
