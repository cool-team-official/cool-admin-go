package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/dict/model"
)

type DictTypeService struct {
	*cool.Service
}

func NewDictTypeService() *DictTypeService {
	return &DictTypeService{
		Service: &cool.Service{
			Model: model.NewDictType(),
			ListQueryOp: &cool.ListQueryOp{
				KeyWorkField: []string{"name"},
			},
		},
	}
}
