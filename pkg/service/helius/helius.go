package helius

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/consolelabs/hackerhouse-demo/pkg/config"
	"github.com/consolelabs/hackerhouse-demo/pkg/logger"
	"github.com/consolelabs/hackerhouse-demo/pkg/model"
	"github.com/consolelabs/hackerhouse-demo/pkg/store"
)

type heliusService struct {
	config *config.Config
	logger logger.Logger
	store  *store.Store
}

func New(cfg *config.Config, l logger.Logger, store *store.Store) IService {
	return &heliusService{
		config: cfg,
		logger: l,
		store:  store,
	}
}

func (h *heliusService) GetMintList(creator string) ([]model.MintData, error) {
	var mintList []model.MintData
	paginationToken := ""
	for {
		response, err := h.getHeliusMintList(creator, paginationToken)
		if err != nil {
			return mintList, err
		}
		if response != nil {
			mintList = append(mintList, response.Result...)
		}
		if response.PaginationToken == "" {
			break
		}
		paginationToken = response.PaginationToken
	}
	return mintList, nil
}

func (h *heliusService) getHeliusMintList(creator, paginationToken string) (response *model.HeliusReponse, err error) {
	retry := 0
	for retry < 5 {
		response, err = h.fetchHeliusMintList(creator, paginationToken)
		if err != nil {
			h.logger.Fields(logger.Fields{
				"creator":         creator,
				"paginationToken": paginationToken,
				"retry":           retry,
			}).Error(err, "failed to get mint list from helius, retrying")
			retry++
			time.Sleep(3 * time.Duration(retry) * time.Second)
			continue
		}
		break
	}

	if retry == 5 {
		return nil, err
	}

	return response, nil
}

func (h *heliusService) fetchHeliusMintList(creator string, paginationToken string) (*model.HeliusReponse, error) {
	var client = &http.Client{}

	requestBody := model.MintListRequest{}
	requestBody.Query.FirstVerifiedCreators = []string{creator}
	requestBody.Options.Limit = 5000
	if paginationToken == "" {
		requestBody.Options.PaginationToken = nil
	} else {
		requestBody.Options.PaginationToken = &paginationToken
	}

	requestBodyByte, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/mintlist?api-key=%s", h.config.Helius.API, h.config.Helius.Key)

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
	res := &model.HeliusReponse{}
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
