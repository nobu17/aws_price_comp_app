package repositories

import (
	"common/aws"
	"common/log"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const afuncName = "Go_Alert"

type alertLambdaParam struct {
	Method   string          `json:"method"`
	GetParam AlertGetRequest `json:"get_param"`
	PutParam AlertPutRequest `json:"put_param"`
}

type alertambdaRepositories struct {
	logger log.LoggerImpl
}

// NewAlertambdaRepositories constructor.
func NewAlertambdaRepositories(logger log.LoggerImpl) AlertImpl {
	return &alertambdaRepositories{logger: logger}
}

// GetAlerts get alert info.
func (u *alertambdaRepositories) GetAlerts(req AlertGetRequest) (AlertGetResponce, error) {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	param := alertLambdaParam{Method: "get", GetParam: req}
	u.logger.LogWriteWithMsgAndObj(log.Info, "lambda input:", param)
	res, err := aws.CallLambdaWithSync(afuncName, region, param)
	if err != nil {
		u.logger.LogWrite(log.Error, "lambda call is failed"+fmt.Sprint(err))
		return AlertGetResponce{}, err
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "lambda result:", *res)
	u.logger.LogWrite(log.Info, "lambda payload:"+string(res.Payload))
	if *res.StatusCode != 200 {
		u.logger.LogWrite(log.Error, "lambda call StatusCode is not 200"+fmt.Sprint(*res.StatusCode))
		return AlertGetResponce{}, errors.New("lambda call StatusCode is not 200" + fmt.Sprint(*res.StatusCode))
	}
	var responce AlertGetResponce
	err = json.Unmarshal(res.Payload, &responce)
	if err != nil {
		u.logger.LogWrite(log.Error, "json Unmarshal is failed"+fmt.Sprint(err))
		return AlertGetResponce{}, err
	}
	return responce, nil
}

// PutAlerts get alert info.
func (u *alertambdaRepositories) PutAlerts(req AlertPutRequest) (AlertPutResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start alertambdaRepositories:PutAlerts", req)
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	param := alertLambdaParam{Method: "put", PutParam: req}
	res, err := aws.CallLambdaWithSync(afuncName, region, param)
	if err != nil {
		u.logger.LogWrite(log.Error, "lambda call is failed"+fmt.Sprint(err))
		return AlertPutResponce{}, err
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "lambda result:", *res)
	u.logger.LogWrite(log.Info, "lambda payload:"+string(res.Payload))
	if *res.StatusCode != 200 {
		u.logger.LogWrite(log.Error, "lambda call StatusCode is not 200"+fmt.Sprint(*res.StatusCode))
		return AlertPutResponce{}, errors.New("lambda call StatusCode is not 200" + fmt.Sprint(*res.StatusCode))
	}
	var responce AlertPutResponce
	err = json.Unmarshal(res.Payload, &responce)
	if err != nil {
		u.logger.LogWrite(log.Error, "json Unmarshal is failed"+fmt.Sprint(err))
		return AlertPutResponce{}, err
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end alertambdaRepositories:PutAlerts", responce)
	return responce, nil
}
