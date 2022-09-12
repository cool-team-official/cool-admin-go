package service

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/dict/model"
	"github.com/gogf/gf/v2/frame/g"
)

type DictInfoService struct {
	*cool.Service
}

// Data方法, 用于获取数据
func (s *DictInfoService) Data(ctx context.Context, types []string) (data interface{}, err error) {
	var (
		// dictInfoModel = model.NewDictInfo()
		dictTypeModel = model.NewDictType()
	)
	mType := cool.GDBM(dictTypeModel)
	// 如果types不为空, 则查询指定类型的数据
	if len(types) > 0 {
		mType = mType.Where("type in (?)", types)
	}
	// 查询所有类型
	typeData, err := mType.All()
	// 如果typeData为空, 则返回空
	if len(typeData) == 0 {
		return g.Map{}, nil
	}

	return
}

// NewDictInfoService 初始化 DictInfoService
func NewDictInfoService() *DictInfoService {
	return &DictInfoService{
		&cool.Service{
			Model: model.NewDictInfo(),
			ListQueryOp: &cool.QueryOp{
				FieldEQ:      []string{"typeId"},
				KeyWordField: []string{"name"},
				AddOrderby:   g.MapStrStr{"createTime": "ASC"},
			},
		},
	}
}
