package repositories

import (
	"common/log"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

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

func (u *dynamoRepository) getTable() dynamo.Table {
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	table := db.Table("SendAlertLog")
	return table
}
