package service

import (
	"github.com/consolelabs/hackerhouse-demo/pkg/config"
	"github.com/consolelabs/hackerhouse-demo/pkg/logger"
	"github.com/consolelabs/hackerhouse-demo/pkg/service/helius"
	mochi_api "github.com/consolelabs/hackerhouse-demo/pkg/service/mochi-api"
	"github.com/consolelabs/hackerhouse-demo/pkg/store"
)

type Service struct {
	Helius helius.IService
	Mochi  mochi_api.IService
}

func New(cfg *config.Config, l logger.Logger, store *store.Store) *Service {
	return &Service{
		Helius: helius.New(cfg, l, store),
		Mochi:  mochi_api.New(cfg, l, store),
	}
}
