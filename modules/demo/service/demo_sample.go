package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/demo/model"
)

type DemoSampleService struct {
	*cool.Service
}

func NewDemoSampleService() *DemoSampleService {
	return &DemoSampleService{
		&cool.Service{
			Model: model.NewDemoSample(),
		},

		// Service: cool.NewService(model.NewDemoSample()),
	}
}
