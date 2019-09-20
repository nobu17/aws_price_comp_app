package repositories

import (
	"common/aws"
	"common/log"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const funcName = "Go_User"

// deafult dynamoregion
const envRegionKey = "REGION"
const defaultRegion = "ap-northeast-1"

type lambdaParam struct {
	Method   string   `json:"method"`
	GetParam userIfno `json:"get_param"`
}

type userIfno struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type lambdaRepositories struct {
	logger log.LoggerImpl
}

// NewLambdaRepositories constructor.
func NewLambdaRepositories(logger log.LoggerImpl) UserInfoImpl {
	return &lambdaRepositories{logger: logger}
}

// GetUserInfo get user info.
func (u *lambdaRepositories) GetUserInfo(req GetRequest) (GetResponce, error) {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	param := lambdaParam{Method: "get", GetParam: userIfno{UserID: req.UserID, Password: req.Password}}
	res, err := aws.CallLambdaWithSync(funcName, region, param)
	if err != nil {
		u.logger.LogWrite(log.Error, "lambda call is failed"+fmt.Sprint(err))
		return GetResponce{}, err
	}
	if *res.StatusCode != 200 {
		u.logger.LogWrite(log.Error, "lambda call StatusCode is not 200"+fmt.Sprint(*res.StatusCode))
		return GetResponce{}, errors.New("lambda call StatusCode is not 200" + fmt.Sprint(*res.StatusCode))
	}
	var responce GetResponce
	err = json.Unmarshal(res.Payload, &responce)
	if err != nil {
		u.logger.LogWrite(log.Error, "json Unmarshal is failed"+fmt.Sprint(err))
		return GetResponce{}, err
	}
	return responce, nil
}
