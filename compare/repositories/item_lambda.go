package repositories

import (
	"common/aws"
	"common/log"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const ifuncName = "Go_Item"

// deafult dynamoregion
const envRegionKey = "REGION"
const defaultRegion = "ap-northeast-1"

type itemGetLambdaParam struct {
	Method   string   `json:"method"`
	GetParam userIfno `json:"get_param"`
}

type userIfno struct {
	UserID  string `json:"user_id"`
	GroupID string `json:"group_id"`
}

type itemLambdaRepositories struct {
	logger log.LoggerImpl
}

// NewItemLambdaRepositories constructor.
func NewItemLambdaRepositories(logger log.LoggerImpl) ItemImpl {
	return &itemLambdaRepositories{logger: logger}
}

// GetUserInfo get user info.
func (u *itemLambdaRepositories) GetItems(req ItemGetRequest) (ItemGetResponce, error) {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	param := itemGetLambdaParam{Method: "get", GetParam: userIfno{UserID: req.UserID, GroupID: req.GroupID}}
	res, err := aws.CallLambdaWithSync(ifuncName, region, param)
	if err != nil {
		u.logger.LogWrite(log.Error, "lambda call is failed"+fmt.Sprint(err))
		return ItemGetResponce{}, err
	}
	if *res.StatusCode != 200 {
		u.logger.LogWrite(log.Error, "lambda call StatusCode is not 200"+fmt.Sprint(*res.StatusCode))
		return ItemGetResponce{}, errors.New("lambda call StatusCode is not 200" + fmt.Sprint(*res.StatusCode))
	}
	var responce ItemGetResponce
	err = json.Unmarshal(res.Payload, &responce)
	if err != nil {
		u.logger.LogWrite(log.Error, "json Unmarshal is failed"+fmt.Sprint(err))
		return ItemGetResponce{}, err
	}
	return responce, nil
}
