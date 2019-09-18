package repositories

import (
	"common/log"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// deafult dynamoregion
const envRegionKey = "REGION"
const defaultRegion = "ap-northeast-1"

// dynamoRepository interface
type dynamoRepository struct {
	logger log.LoggerImpl
}

// NewDynamoRepository constructor
func NewDynamoRepository(logger log.LoggerImpl) AlertImpl {
	return &dynamoRepository{logger: logger}
}

// sendAlertLog struct
type sendAlertLog struct {
	// ユーザーID
	UserID string `dynamo:"user_id"`
	// アラート日付
	AlertDate string `dynamo:"alert_date"`
	// 店鋪種類
	StoreType string `dynamo:"store_type"`
	// 商品ID
	ProductID string `dynamo:"product_id"`
	// 価格
	Price int `dynamo:"price"`
}

// GetAlertLog impl
func (u *dynamoRepository) GetAlertLog(req GetRequest) (GetResponce, error) {
	var items []sendAlertLog
	table := u.getTable()
	err := table.Get("user_id", req.UserID).Range("alert_date", dynamo.GreaterOrEqual, req.MinAlertDate).All(&items)
	if err != nil {
		u.logger.LogWrite(log.Error, "table.Get Error"+fmt.Sprint(err))
		return GetResponce{}, err
	}
	var list = make([]SendAlertLog, 0)
	for _, item := range items {
		var t = NewSendAlertLog(item.UserID, item.AlertDate, item.StoreType, item.ProductID, item.Price)
		list = append(list, t)
	}
	return GetResponce{SendAlertLogList: list}, nil
}

func (u *dynamoRepository) PutAlertLog(req PutRequest) (PutResponce, error) {
	u.logger.LogWrite(log.Info, "start PutAlertLog")
	table := u.getTable()
	batch := table.Batch().Write().Put(req.PutAlertLogList)
	wrote, err := batch.Run()
	if err != nil {
		u.logger.LogWrite(log.Error, "error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutAlertLog")
		return PutResponce{Wrote: wrote}, err
	}
	if len(req.PutAlertLogList) != wrote {
		u.logger.LogWrite(log.Error, fmt.Sprintf("some wrote is failed. Total:%v Success:%v", len(req.PutAlertLogList), wrote))
	}
	u.logger.LogWrite(log.Info, "end PutAlertLog")
	return PutResponce{Wrote: wrote}, err
}

func (u *dynamoRepository) getTable() dynamo.Table {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String(region),
	})
	table := db.Table("SendAlertLog")
	return table
}
