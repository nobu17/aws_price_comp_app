package repositories

import (
	"common/aws"
	"common/log"
	"fmt"
	"os"
)

const ifuncName = "Go_Item"

type itemGroupDleteLambdaParam struct {
	Method      string   `json:"method"`
	DeleteParam userIfno `json:"delete_param"`
}

type userIfno struct {
	UserID  string `json:"user_id"`
	GroupID string `json:"group_id"`
}

type itemGroupLambdaRepositories struct {
	logger log.LoggerImpl
}

// NewItemGroupLambdaRepositories constructor.
func NewItemGroupLambdaRepositories(logger log.LoggerImpl) GroupInfoImpl {
	return &itemGroupLambdaRepositories{logger: logger}
}

// DeleteItemGroup get user info.
func (u *itemGroupLambdaRepositories) DeleteItemGroup(req DeleteItemGroupRequest) (DeleteItemGroupResponce, error) {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}
	var successGroups = make([]string, 0)
	var faileGroups = make([]string, 0)

	for _, group := range req.GroupIDList {
		param := itemGroupDleteLambdaParam{Method: "delete", DeleteParam: userIfno{UserID: req.UserID, GroupID: group}}
		res, err := aws.CallLambdaWithSync(ifuncName, region, param)
		if err != nil {
			u.logger.LogWrite(log.Error, "lambda call is failed"+fmt.Sprint(err))
			faileGroups = append(faileGroups, group)
			continue
		}
		if *res.StatusCode != 200 {
			u.logger.LogWrite(log.Error, "lambda call StatusCode is not 200"+fmt.Sprint(*res.StatusCode))
			faileGroups = append(faileGroups, group)
			continue
		}
		successGroups = append(successGroups, group)
	}

	return DeleteItemGroupResponce{SuccessItemGroupList: successGroups, FailedItemGroupList: faileGroups}, nil
}
