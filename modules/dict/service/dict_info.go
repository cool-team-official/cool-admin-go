package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/dict/model"
)

type DictInfoService struct {
	*cool.Service
}

func NewDictInfoService() *DictInfoService {
	return &DictInfoService{
		Service: cool.NewService(model.NewDictInfo()),
	}
}
