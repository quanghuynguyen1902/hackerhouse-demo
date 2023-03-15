package model

type HeliusReponse struct {
	Result          []MintData `json:"result"`
	PaginationToken string     `json:"paginationToken"`
}

type MintData struct {
	Mint string `json:"mint"`
	Name string `json:"name"`
}

type MintListRequest struct {
	Query struct {
		FirstVerifiedCreators []string `json:"firstVerifiedCreators"`
	} `json:"query"`
	Options struct {
		Limit           int     `json:"limit"`
		PaginationToken *string `json:"paginationToken"`
	} `json:"options"`
}
