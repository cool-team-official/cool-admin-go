package service

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type BaseSysParamService struct {
	*cool.Service
}

func NewBaseSysParamService() *BaseSysParamService {
	return &BaseSysParamService{
		&cool.Service{
			Model: model.NewBaseSysParam(),
		},

		// Service: cool.NewService(model.NewBaseSysParam()),
	}
}

// HtmlByKey 根据配置参数key获取网页内容(富文本)
func (s *BaseSysParamService) HtmlByKey(key string) string {
	var (
		html = "<html><body>@content</body></html>"
	)
	m := cool.DBM(s.Model)
	record, err := m.Where("keyName = ?", key).One()
	if err != nil {
		html = gstr.Replace(html, "@content", err.Error())
		return html
	}
	if record.IsEmpty() {
		html = gstr.Replace(html, "@content", "key notfound")
		return html
	}
	html = gstr.Replace(html, "@content", record["data"].String())

	return html
}

// ModifyAfter 修改后
func (s *BaseSysParamService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	var (
		m = cool.DBM(s.Model)
	)
	result, err := m.All()
	if err != nil {
		return
	}
	for _, v := range result {
		key := "param:" + v["keyName"].String()
		value := v["data"].String()
		err = cool.CacheManager.Set(ctx, key, value, 0)
		if err != nil {
			return
		}
	}
	return
}

// DataByKey 根据配置参数key获取数据
func (s *BaseSysParamService) DataByKey(ctx context.Context, key string) (data string, err error) {
	var (
		m = cool.DBM(s.Model)
	)
	rKey := "param:" + key
	dataCache, err := cool.CacheManager.Get(ctx, rKey)
	if err != nil {
		return
	}
	if !dataCache.IsEmpty() {
		data = dataCache.String()
		return
	}
	record, err := m.Where("keyName = ?", key).One()
	if err != nil {
		return
	}
	if record.IsEmpty() {
		return
	}
	data = record["data"].String()
	err = cool.CacheManager.Set(ctx, rKey, data, 0)
	return
}
