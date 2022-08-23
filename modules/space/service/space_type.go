package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/space/model"
)

type SpaceTypeService struct {
	*cool.Service
}

func NewSpaceTypeService() *SpaceTypeService {
	return &SpaceTypeService{
		&cool.Service{
			Model: model.NewSpaceType(),
		},

		// Service: cool.NewService(model.NewSpaceType()),
	}
}
