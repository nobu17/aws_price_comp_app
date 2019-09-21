package repositories

import (
	"common/aws"
	"common/log"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const nfuncName = "Go_Notify"

type notifyLambdaParam struct {
	Method   string           `json:"method"`
	PutParam NotifyPutRequest `json:"put_param"`
}

type notifyLambdaRepositories struct {
	logger log.LoggerImpl
}

// NewNotifyLambdaRepositories constructor.
func NewNotifyLambdaRepositories(logger log.LoggerImpl) NotifyImpl {
	return &notifyLambdaRepositories{logger: logger}
}

// PutAlerts get alert info.
func (u *notifyLambdaRepositories) SendNotify(req NotifyPutRequest) (NotifyPutResponce, error) {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	param := notifyLambdaParam{Method: "put", PutParam: req}
	res, err := aws.CallLambdaWithSync(nfuncName, region, param)
	if err != nil {
		u.logger.LogWrite(log.Error, "lambda call is failed"+fmt.Sprint(err))
		return NotifyPutResponce{}, err
	}
	if *res.StatusCode != 200 {
		u.logger.LogWrite(log.Error, "lambda call StatusCode is not 200"+fmt.Sprint(*res.StatusCode))
		return NotifyPutResponce{}, errors.New("lambda call StatusCode is not 200" + fmt.Sprint(*res.StatusCode))
	}
	var responce NotifyPutResponce
	err = json.Unmarshal(res.Payload, &responce)
	if err != nil {
		u.logger.LogWrite(log.Error, "json Unmarshal is failed"+fmt.Sprint(err))
		return NotifyPutResponce{}, err
	}
	return responce, nil
}
