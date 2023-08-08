package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/v2/net/gsvc"
)

type DemoTestService struct {
	*cool.Service
}

func NewDemoTestService() *DemoTestService {
	return &DemoTestService{
		&cool.Service{},
	}
}

func (s *DemoTestService) GetDemoTestList() (interface{}, error) {
	gsvc.SetRegistry(etcd.New(`127.0.0.1:2379`))
	return nil, nil
}
