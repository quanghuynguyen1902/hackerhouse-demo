package nft

import (
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) INft {
	return &store{
		db: db,
	}
}

func (s *store) GetLatestTokenId(address string) (int, error) {
	var latestTokenId int
	row := s.db.Table("nft_token").Select("max(CAST(token_id as INT))").Where("collection_address = ?", address).Row()
	err := row.Scan(&latestTokenId)
	if err != nil {
		return 0, err
	}
	return latestTokenId, nil
}
