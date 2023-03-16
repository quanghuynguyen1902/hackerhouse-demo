package helius

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/consolelabs/hackerhouse-demo/pkg/config"
	"github.com/consolelabs/hackerhouse-demo/pkg/logger"
	"github.com/consolelabs/hackerhouse-demo/pkg/model"
	mochi_api "github.com/consolelabs/hackerhouse-demo/pkg/service/mochi-api"
	"io/ioutil"
	"net/http"
)

type heliusService struct {
	config *config.Config
	logger logger.Logger
}

func New(cfg *config.Config, l logger.Logger) IService {
	return &heliusService{
		config: cfg,
		logger: l,
	}
}

func (h *heliusService) GetNftFromTransaction(signature string) (interface{}, error) {
	mochi := mochi_api.New(h.config, h.logger)
	parseTransactionData, err := h.parseTransactionWithHelius(signature)
	if err != nil {
		return nil, err
	}
	if len(parseTransactionData) > 0 {
		mintListData := parseTransactionData[0].Events.Nft.Nfts
		mintList := []string{}
		for _, e := range mintListData {
			mintList = append(mintList, e.Mint)
		}
		tokenMetaList, err := h.getTokenMetadataWithHelius(mintList)
		if err != nil {
			logger.L.Fields(logger.Fields{}).Error(err, "failed get data from token metadata")
			return nil, err
		}
		for _, t := range tokenMetaList {
			mochiData, err := mochi.GetNftDetail(t.OnChainMetadata.Metadata.Data.Name)
			if err != nil {
				logger.L.Fields(logger.Fields{
					"token": t.Account,
				}).Error(err, "failed get data from mochi")
			}
			fmt.Println(mochiData)
		}
	}

	return nil, nil
}

func (h *heliusService) parseTransactionWithHelius(signature string) ([]model.HeliusTransactionData, error) {
	var client = &http.Client{}

	requestBody := model.ParseTransactionRequest{}

	requestBody.Transactions = []string{signature}

	requestBodyByte, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v0/transactions/?api-key=%s", h.config.Helius.API, h.config.Helius.Key)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyByte))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var res []model.HeliusTransactionData
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *heliusService) getTokenMetadataWithHelius(mintList []string) ([]model.HeliusTokenMetadata, error) {
	var client = &http.Client{}

	requestBody := model.HeliusTokenMetadataRequest{}

	requestBody.MintAccounts = mintList
	requestBody.IncludeOffChain = false

	requestBodyByte, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v0/token-metadata/?api-key=%s", h.config.Helius.API, h.config.Helius.Key)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyByte))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var res []model.HeliusTokenMetadata
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
