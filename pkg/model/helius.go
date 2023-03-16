package model

type ParseTransactionRequest struct {
	Transactions []string `json:"transactions"`
}

type HeliusTransactionData struct {
	Description     string        `json:"description"`
	Type            string        `json:"type"`
	Source          string        `json:"source"`
	Fee             int           `json:"fee"`
	FeePayer        string        `json:"feePayer"`
	Signature       string        `json:"signature"`
	Slot            int           `json:"slot"`
	Timestamp       int           `json:"timestamp"`
	TokenTransfers  []interface{} `json:"tokenTransfers"`
	NativeTransfers []interface{} `json:"nativeTransfers"`
	AccountData     []struct {
		Account             string        `json:"account"`
		NativeBalanceChange int           `json:"nativeBalanceChange"`
		TokenBalanceChanges []interface{} `json:"tokenBalanceChanges"`
	} `json:"accountData"`
	TransactionError interface{} `json:"transactionError"`
	Instructions     []struct {
		Accounts          []string      `json:"accounts"`
		Data              string        `json:"data"`
		ProgramID         string        `json:"programId"`
		InnerInstructions []interface{} `json:"innerInstructions"`
	} `json:"instructions"`
	Events struct {
		Nft struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Source      string `json:"source"`
			Amount      int64  `json:"amount"`
			Fee         int    `json:"fee"`
			FeePayer    string `json:"feePayer"`
			Signature   string `json:"signature"`
			Slot        int    `json:"slot"`
			Timestamp   int    `json:"timestamp"`
			SaleType    string `json:"saleType"`
			Buyer       string `json:"buyer"`
			Seller      string `json:"seller"`
			Staker      string `json:"staker"`
			Nfts        []struct {
				Mint          string `json:"mint"`
				TokenStandard string `json:"tokenStandard"`
			} `json:"nfts"`
		} `json:"nft"`
	} `json:"events"`
}

type HeliusTokenMetadataRequest struct {
	MintAccounts    []string `json:"mintAccounts"`
	IncludeOffChain bool     `json:"includeOffChain"`
}

type HeliusTokenMetadata struct {
	Account            string `json:"account"`
	OnChainAccountInfo struct {
		AccountInfo struct {
			Key        string `json:"key"`
			IsSigner   bool   `json:"isSigner"`
			IsWritable bool   `json:"isWritable"`
			Lamports   int    `json:"lamports"`
			Data       struct {
				Parsed struct {
					Info struct {
						Decimals        int    `json:"decimals"`
						FreezeAuthority string `json:"freezeAuthority"`
						IsInitialized   bool   `json:"isInitialized"`
						MintAuthority   string `json:"mintAuthority"`
						Supply          string `json:"supply"`
					} `json:"info"`
					Type string `json:"type"`
				} `json:"parsed"`
				Program string `json:"program"`
				Space   int    `json:"space"`
			} `json:"data"`
			Owner      string `json:"owner"`
			Executable bool   `json:"executable"`
			RentEpoch  int    `json:"rentEpoch"`
		} `json:"accountInfo"`
		Error string `json:"error"`
	} `json:"onChainAccountInfo"`
	OnChainMetadata struct {
		Metadata struct {
			TokenStandard   string `json:"tokenStandard"`
			Key             string `json:"key"`
			UpdateAuthority string `json:"updateAuthority"`
			Mint            string `json:"mint"`
			Data            struct {
				Name                 string `json:"name"`
				Symbol               string `json:"symbol"`
				URI                  string `json:"uri"`
				SellerFeeBasisPoints int    `json:"sellerFeeBasisPoints"`
				Creators             []struct {
					Address  string `json:"address"`
					Verified bool   `json:"verified"`
					Share    int    `json:"share"`
				} `json:"creators"`
			} `json:"data"`
			PrimarySaleHappened bool `json:"primarySaleHappened"`
			IsMutable           bool `json:"isMutable"`
			EditionNonce        int  `json:"editionNonce"`
			Uses                struct {
				UseMethod string `json:"useMethod"`
				Remaining int    `json:"remaining"`
				Total     int    `json:"total"`
			} `json:"uses"`
			Collection struct {
				Key      string `json:"key"`
				Verified bool   `json:"verified"`
			} `json:"collection"`
			CollectionDetails interface{} `json:"collectionDetails"`
		} `json:"metadata"`
		Error string `json:"error"`
	} `json:"onChainMetadata"`
	LegacyMetadata interface{} `json:"legacyMetadata"`
}
