package main

import (
	"fmt"
	"github.com/consolelabs/hackerhouse-demo/pkg/config"
	"github.com/consolelabs/hackerhouse-demo/pkg/logger"
	"github.com/consolelabs/hackerhouse-demo/pkg/service"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	l := logger.NewLogrusLogger()

	svc := service.New(cfg, l)
	// y00t collection
	mintList, err := svc.Helius.GetNftFromTransaction("5SiXMSK9E4SwAZt7XawCdg1yRa6og5VBnyZnZYnkErrCbCbbdv5hGkKXGUFex59nNqLHrfA4q61JvEDdYeP8gkgU")
	if err != nil {
	}
	fmt.Println(mintList)
	//

}
