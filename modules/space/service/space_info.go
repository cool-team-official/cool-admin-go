package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/space/model"
)

type SpaceInfoService struct {
	*cool.Service
}

func NewSpaceInfoService() *SpaceInfoService {
	return &SpaceInfoService{
		&cool.Service{
			Model: model.NewSpaceInfo(),
		},

		// Service: cool.NewService(model.NewSpaceInfo()),
	}
}
