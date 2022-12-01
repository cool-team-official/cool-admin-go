package service

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/dict/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type DictInfoService struct {
	*cool.Service
}

// Data方法, 用于获取数据
func (s *DictInfoService) Data(ctx context.Context, types []string) (data interface{}, err error) {
	var (
		dictInfoModel = model.NewDictInfo()
		dictTypeModel = model.NewDictType()
	)
	mType := cool.DBM(dictTypeModel)
	// 如果types不为空, 则查询指定类型的数据
	if len(types) > 0 {
		mType = mType.Where("type in (?)", types)
	}
	// 查询所有类型
	typeData, err := mType.All()
	// 如果typeData为空, 则返回空
	if typeData.IsEmpty() {
		return g.Map{}, nil
	}
	data = g.Map{}
	for _, v := range typeData {
		m := cool.DBM(dictInfoModel)
		result, err := m.Where("typeId=?", v["id"]).Fields("id", "name", "parentId", "typeId").Order("orderNum asc").All()
		if err != nil {
			return nil, err
		}
		if result.IsEmpty() {
			continue
		}
		data.(g.Map)[v["key"].String()] = result
	}
	return
}

// ModifyAfter 修改后
func (s *DictInfoService) ModifyAfter(ctx context.Context, method string, param map[string]interface{}) (err error) {
	if method == "Delete" {
		// 删除后,同时删除子节点
		ids, ok := param["ids"]
		if !ok {
			return
		}
		for _, v := range ids.([]interface{}) {
			err = delChildDict(gconv.Int64(v))
			if err != nil {
				return
			}
		}
	}
	return
}

// delChildDict 删除子字典
func delChildDict(id int64) error {
	var (
		dictInfoModel = model.NewDictInfo()
	)
	m := cool.DBM(dictInfoModel)
	result, err := m.Where("parentId=?", id).Fields("id").All()
	if err != nil {
		return err
	}
	if result.IsEmpty() {
		return nil
	}
	for _, v := range result {
		delChildDict(v["id"].Int64())
	}
	_, err = m.Where("parentId=?", id).Delete()
	return err
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
