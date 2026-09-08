package service

import (
	"strings"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/encoding/gjson"
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
				KeyWordField: []string{"name", "params", "ipAddr"},
				Select:       "base_sys_log.*,user.name ",
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
	// 请求体含二进制内容(文件上传等)时,读取整段会放大内存且无检索价值,
	// 一律不读取body,仅记录占位信息
	contentType := r.Header.Get("Content-Type")
	switch {
	case strings.Contains(contentType, "multipart/"):
		baseSysLog.Params = "[文件上传内容省略]"
	case contentType != "" &&
		!strings.Contains(contentType, "json") &&
		!strings.Contains(contentType, "text/") &&
		!strings.Contains(contentType, "x-www-form-urlencoded"):
		baseSysLog.Params = "[二进制内容省略]"
	default:
		// 其余请求体脱敏(剔除密码/验证码/令牌等敏感字段)并限制长度后再落库
		baseSysLog.Params = sanitizeLogParams(r.GetBodyString())
	}
	m := cool.DBM(s.Model)
	if _, err := m.Insert(g.Map{
		"userId": baseSysLog.UserID,
		"action": baseSysLog.Action,
		"ip":     baseSysLog.IP,
		"ipAddr": baseSysLog.IPAddr,
		"params": baseSysLog.Params,
	}); err != nil {
		g.Log().Error(ctx, "操作日志写入失败", err)
	}
}

// maxLogParamsLen 单条日志参数最大保留长度,防止超大请求体写入日志表
const maxLogParamsLen = 4000

// truncateLogParams 超长内容截断,保证日志写入稳定
func truncateLogParams(s string) string {
	if len(s) > maxLogParamsLen {
		return s[:maxLogParamsLen] + "...[内容过长已截断]"
	}
	return s
}

// 请求体中可能出现的敏感字段,落库前需剔除/隐藏
var logSensitiveKeys = []string{
	"password",
	"oldPassword",
	"newPassword",
	"verifyCode",
	"refreshToken",
	"accessToken",
	"secret",
}

// sanitizeLogParams 移除请求体中的敏感字段并限制长度后返回,避免明文密码/验证码/令牌等写入日志
func sanitizeLogParams(body string) string {
	if body == "" {
		return body
	}
	if json, err := gjson.LoadContent([]byte(body)); err == nil {
		for _, key := range logSensitiveKeys {
			_ = json.Remove(key)
		}
		if s, err := json.ToJsonString(); err == nil {
			return truncateLogParams(s)
		}
		return "[敏感参数已隐藏]"
	}
	// 非标准JSON(如表单/查询串),若包含敏感字段则整体隐藏
	lower := strings.ToLower(body)
	for _, key := range logSensitiveKeys {
		if strings.Contains(lower, strings.ToLower(key)) {
			return "[敏感参数已隐藏]"
		}
	}
	return truncateLogParams(body)
}

// Clear 清除日志
func (s *BaseSysLogService) Clear(isAll bool) (err error) {
	BaseSysConfService := NewBaseSysConfService()
	m := cool.DBM(s.Model)
	if isAll {
		_, err = m.Delete("1=1")
	} else {
		keepDays := gconv.Int(BaseSysConfService.GetValue("logKeep"))
		_, err = m.Delete("createTime < ?", gtime.Now().AddDate(0, 0, -keepDays).String())
	}
	return
}
