package helius

type IService interface {
	GetNftFromTransaction(signature string) (interface{}, error)
}
