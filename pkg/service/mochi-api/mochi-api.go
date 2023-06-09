package mochi_api

import (
	"encoding/json"
	"fmt"
	"github.com/consolelabs/hackerhouse-demo/pkg/config"
	"github.com/consolelabs/hackerhouse-demo/pkg/logger"
	"io/ioutil"
	"net/http"
	"strings"
)

type mochiService struct {
	config *config.Config
	logger logger.Logger
}

func New(cfg *config.Config, l logger.Logger) IService {
	return &mochiService{
		config: cfg,
		logger: l,
	}
}

func (m *mochiService) GetNftDetail(name string) (interface{}, error) {
	tokenId := ""
	if strings.Contains(name, "#") {
		tokenId = strings.Split(name, "#")[1]
	}
	nftToken, err := m.fetchNftDetail(tokenId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return nftToken, nil
}

func (m *mochiService) fetchNftDetail(tokenId string) (interface{}, error) {
	collectionAddress := "solscan-229f30fb8b5f0a7ff7fea1acd51bd102be43fe02e8d1c24f36331b41dae0d167"
	var client = &http.Client{}

	url := fmt.Sprintf("https://api.indexer.console.so/api/v1/nft/%s/%s", collectionAddress, tokenId)

	request, err := http.NewRequest("GET", url, nil)
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
	var res interface{}
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
