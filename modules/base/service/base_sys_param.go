package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
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
