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

type priceLogMaster struct {
	// ユーザーID
	UserID string `dynamo:"user_id"`
	// グループID
	GroupID string `dynamo:"group_id"`
	// 価格リスト
	PriceList []priceLog `dynamo:"price_list"`
}

type priceLog struct {
	StoreType            string `dynamo:"store_type"`
	ItemID               string `dynamo:"item_id"`
	Price                int    `dynamo:"price"`
	LastModifiedDatetime string `dynamo:"last_modified_datetime"`
}

// NewDynamoRepository constructor
func NewDynamoRepository(logger log.LoggerImpl) PriceLogImpl {
	return &dynamoRepository{logger: logger}
}

func (u *dynamoRepository) GetPriceLogs(req GetRequest) (GetResponce, error) {
	var pLogs []priceLogMaster
	table := u.getPriceLog()
	err := table.Get("user_id", req.UserID).Filter("group_id = ?", req.GroupID).All(&pLogs)
	if err != nil {
		u.logger.LogWrite(log.Error, "priceLogs.Get Error"+fmt.Sprint(err))
		return GetResponce{}, err
	}

	if len(pLogs) < 1 {
		u.logger.LogWrite(log.Error, "no data")
	}

	var list = make([]PriceLog, 0)
	for _, price := range pLogs[0].PriceList {
		var item = PriceLog{StoreType: price.StoreType, ItemID: price.ItemID, Price: price.Price, LastModifiedDatetime: price.LastModifiedDatetime}
		list = append(list, item)
	}
	responce := GetResponce{UserID: pLogs[0].UserID, GroupID: pLogs[0].GroupID, PriceLogList: list}

	return responce, nil
}

func (u *dynamoRepository) PutPriceLogs(req PutRequest) (PutResponce, error) {
	var list = make([]priceLog, 0)
	for _, price := range req.PriceLogList {
		var item = priceLog{StoreType: price.StoreType, ItemID: price.ItemID, Price: price.Price, LastModifiedDatetime: price.LastModifiedDatetime}
		list = append(list, item)
	}
	param := priceLogMaster{UserID: req.UserID, GroupID: req.GroupID, PriceList: list}

	table := u.getPriceLog()
	err := table.Put(param).Run()
	if err != nil {
		u.logger.LogWrite(log.Error, "table.Put Error"+fmt.Sprint(err))
		return PutResponce{}, err
	}

	return PutResponce{}, nil
}

func (u *dynamoRepository) DeletePriceLogs(req DeleteRequest) (DeleteResponce, error) {
	table := u.getPriceLog()
	err := table.Delete("user_id", req.UserID).Range("group_id", req.GroupID).Run()
	if err != nil {
		u.logger.LogWrite(log.Error, "table.Del Error"+fmt.Sprint(err))
		return DeleteResponce{}, err
	}

	return DeleteResponce{}, nil
}

func (u *dynamoRepository) getPriceLog() dynamo.Table {
	db := u.getDB()
	table := db.Table("PriceLogMaster")
	return table
}

func (u *dynamoRepository) getDB() *dynamo.DB {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String(region),
	})
	return db
}
