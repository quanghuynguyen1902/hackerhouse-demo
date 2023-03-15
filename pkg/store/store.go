package store

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/consolelabs/hackerhouse-demo/pkg/config"
	"github.com/consolelabs/hackerhouse-demo/pkg/logger"
	"github.com/consolelabs/hackerhouse-demo/pkg/store/nft"
)

type Store struct {
	Nft nft.INft
}

func New(cfg *config.Config) *Store {
	db := connDb(cfg)
	return &Store{
		Nft: nft.New(db),
	}
}

func connDb(cfg *config.Config) *gorm.DB {
	ds := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Pass,
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Name,
	)

	conn, err := sql.Open("postgres", ds)
	if err != nil {
		logger.L.Fatalf(err, "failed to open database connection")
	}

	db, err := gorm.Open(postgres.New(
		postgres.Config{Conn: conn}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	if err != nil {
		logger.L.Fatalf(err, "failed to open database connection")
	}

	logger.L.Info("database connected")

	return db
}
