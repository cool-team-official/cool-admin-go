package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
)

type Sample struct {
	*cool.Service
}

func NewSample(model interface{}) *Sample {
	return &Sample{
		Service: cool.NewService(model),
	}
}
