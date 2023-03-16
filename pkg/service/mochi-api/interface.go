package mochi_api

type IService interface {
	GetNftDetail(name string) (interface{}, error)
}
