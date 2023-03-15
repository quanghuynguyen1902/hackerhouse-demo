package main

import (
	"fmt"
	"github.com/consolelabs/hackerhouse-demo/pkg/config"
	"github.com/consolelabs/hackerhouse-demo/pkg/logger"
	"github.com/consolelabs/hackerhouse-demo/pkg/service"
	"github.com/consolelabs/hackerhouse-demo/pkg/store"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	l := logger.NewLogrusLogger()
	store := store.New(cfg)

	svc := service.New(cfg, l, store)
	// y00t collection
	mintList, err := svc.Helius.GetMintList("A4FM6h8T5Fmh9z2g3fKUrKfZn6BNFEgByR8QGpdbQhk1")
	if err != nil {
	}

	nftToken, err := svc.Mochi.GetNftDetail(mintList[0])
	fmt.Println(nftToken)

}
