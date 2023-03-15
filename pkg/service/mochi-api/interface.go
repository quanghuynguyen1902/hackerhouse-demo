package mochi_api

import "github.com/consolelabs/hackerhouse-demo/pkg/model"

type IService interface {
	GetNftDetail(mintData model.MintData) (interface{}, error)
}
