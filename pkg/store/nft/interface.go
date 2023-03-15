package nft

type INft interface {
	GetLatestTokenId(address string) (int, error)
}
