package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/demo/model"
)

type DemoGoodsService struct {
	*cool.Service
}

func NewDemoGoodsService() *DemoGoodsService {
	return &DemoGoodsService{
		&cool.Service{
			Model: model.NewDemoGoods(),
			ListQueryOp: &cool.QueryOp{

				Join: []*cool.JoinOp{},
			},
		},
	}
}
