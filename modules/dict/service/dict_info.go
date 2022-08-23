package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/dict/model"
	"github.com/gogf/gf/v2/frame/g"
)

type DictInfoService struct {
	*cool.Service
}

// NewDictInfoService 初始化 DictInfoService
func NewDictInfoService() *DictInfoService {
	return &DictInfoService{
		&cool.Service{
			Model: model.NewDictInfo(),
			ListQueryOp: &cool.ListQueryOp{
				FieldEQ:      []string{"typeId"},
				KeyWorkField: []string{"name"},
				AddOrderby:   g.MapStrStr{"createTime": "ASC"},
			},
		},
	}
}
