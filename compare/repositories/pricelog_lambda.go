package repositories

import (
	"common/aws"
	"common/log"
	"errors"
	"fmt"
	"os"
)

const plfuncName = "Go_PriceLog"

type pricelogLambdaParam struct {
	Method   string             `json:"method"`
	PutParam PutPriceLogRequest `json:"put_param"`
}

type pricelogLambdaRepositories struct {
	logger log.LoggerImpl
}

// NewPricelogLambdaRepositories constructor.
func NewPricelogLambdaRepositories(logger log.LoggerImpl) PriceLogImpl {
	return &pricelogLambdaRepositories{logger: logger}
}

// UpdatePriceLog update pricelogs
func (u *pricelogLambdaRepositories) UpdatePriceLog(req PutPriceLogRequest) (PutPriceLogResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start pricelogLambdaRepositories:UpdatePriceLog", req)
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	param := pricelogLambdaParam{Method: "put", PutParam: req}
	res, err := aws.CallLambdaWithSync(plfuncName, region, param)
	if err != nil {
		u.logger.LogWrite(log.Error, "lambda call is failed"+fmt.Sprint(err))
		return PutPriceLogResponce{}, err
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "lambda result:", *res)
	u.logger.LogWrite(log.Info, "lambda payload:"+string(res.Payload))
	if *res.StatusCode != 200 {
		u.logger.LogWrite(log.Error, "lambda call StatusCode is not 200"+fmt.Sprint(*res.StatusCode))
		return PutPriceLogResponce{}, errors.New("lambda call StatusCode is not 200" + fmt.Sprint(*res.StatusCode))
	}
	var responce = PutPriceLogResponce{}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end alertambdaRepositories:PutAlerts", responce)
	return responce, nil
}
