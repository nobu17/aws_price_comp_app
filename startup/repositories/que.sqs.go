package repositories

import (
	"common/aws"
	"common/info"
	"common/log"
	"encoding/json"
	"fmt"
	"os"
)

type sqsRepositories struct {
	logger log.LoggerImpl
}

// NewSQSRepositories constructor.
func NewSQSRepositories(logger log.LoggerImpl) SendObservImpl {
	return &sqsRepositories{logger: logger}
}

func (u *sqsRepositories) SendObservRequest(req SendRequest) error {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}
	var err error
	for _, itemG := range req.ItemGroupList {
		jsonBytes, err := json.Marshal(itemG)
		if err != nil {
			u.logger.LogWrite(log.Error, "json marshall is failed:"+fmt.Sprint(err))
			break
		}
		messageID, err := aws.SendMessageToSQS(info.GetItemObservSQSURL(), string(jsonBytes), region)
		if err != nil {
			u.logger.LogWrite(log.Error, "send sqs is failed:"+fmt.Sprint(err))
			break
		}
		u.logger.LogWrite(log.Info, "send sqs is sucess:"+messageID)
	}

	return err
}
