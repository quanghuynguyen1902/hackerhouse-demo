package helius

import "github.com/consolelabs/hackerhouse-demo/pkg/model"

type IService interface {
	GetMintList(creator string) ([]model.MintData, error)
}
