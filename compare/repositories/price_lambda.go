package repositories

import (
	"common/aws"
	"common/log"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const pfuncName = "Go_PriceGet"

type priceLambdaRepositories struct {
	logger log.LoggerImpl
}

// NewPriceLambdaRepositories constructor.
func NewPriceLambdaRepositories(logger log.LoggerImpl) PriceImpl {
	return &priceLambdaRepositories{logger: logger}
}

// GetUserInfo get user info.
func (u *priceLambdaRepositories) GetPrices(req PriceGetRequest) (PriceGetResponce, error) {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	res, err := aws.CallLambdaWithSync(pfuncName, region, req)
	if err != nil {
		u.logger.LogWrite(log.Error, "lambda call is failed"+fmt.Sprint(err))
		return PriceGetResponce{}, err
	}
	if *res.StatusCode != 200 {
		u.logger.LogWrite(log.Error, "lambda call StatusCode is not 200"+fmt.Sprint(*res.StatusCode))
		return PriceGetResponce{}, errors.New("lambda call StatusCode is not 200" + fmt.Sprint(*res.StatusCode))
	}
	var responce PriceGetResponce
	err = json.Unmarshal(res.Payload, &responce)
	if err != nil {
		u.logger.LogWrite(log.Error, "json Unmarshal is failed"+fmt.Sprint(err))
		return PriceGetResponce{}, err
	}
	return responce, nil
}
