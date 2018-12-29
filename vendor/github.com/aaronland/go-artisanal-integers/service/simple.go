package service

import (
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aaronland/go-artisanal-integers/utils"
)

type SimpleService struct {
	artisanalinteger.Service
	engine artisanalinteger.Engine
}

func NewSimpleService(eng artisanalinteger.Engine) (*SimpleService, error) {

	svc := SimpleService{
		engine: eng,
	}

	return &svc, nil
}

func (svc *SimpleService) NextInt() (int64, error) {

	i, err := svc.engine.NextInt()

	if err != nil {
		return -1, err
	}

	if utils.IsLondonInteger(i) {
		return svc.NextInt()
	}

	return i, nil
}

func (svc *SimpleService) LastInt() (int64, error) {
	return svc.engine.LastInt()
}
