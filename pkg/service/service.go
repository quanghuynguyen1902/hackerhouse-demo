package service

import (
	"github.com/consolelabs/hackerhouse-demo/pkg/config"
	"github.com/consolelabs/hackerhouse-demo/pkg/logger"
	"github.com/consolelabs/hackerhouse-demo/pkg/service/helius"
	mochi_api "github.com/consolelabs/hackerhouse-demo/pkg/service/mochi-api"
)

type Service struct {
	Helius helius.IService
	Mochi  mochi_api.IService
}

func New(cfg *config.Config, l logger.Logger) *Service {
	return &Service{
		Helius: helius.New(cfg, l),
		Mochi:  mochi_api.New(cfg, l),
	}
}
