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
func NewDynamoRepository(logger log.LoggerImpl) ItemMasterImpl {
	return &dynamoRepository{logger: logger}
}

// itemMaster リクエストの出力パラメータ
type itemMaster struct {
	// ユーザID
	UserID string `dynamo:"user_id"`
	// グループID
	GroupID string `dynamo:"group_id"`
	// 商品ID
	ProductID string `dynamo:"product_id"`
	// 店鋪種類
	StoreType string `dynamo:"store_type"`
	// しきい値
	ThretholdPrice int `dynamo:"threthold_price"`
	// アイテム名
	ItemName string `dynamo:"item_name"`
	// アイテム名
	UniqueID string `dynamo:"unique_id"`
}

// NewitemMaster construtor
func NewitemMaster(userID string, groupID string, productID string, storeType string, thretholdPrice int, itemName string) itemMaster {
	return itemMaster{
		UserID:         userID,
		GroupID:        groupID,
		ProductID:      productID,
		StoreType:      storeType,
		ThretholdPrice: thretholdPrice,
		ItemName:       itemName,
		UniqueID:       storeType + "_" + productID,
	}
}

// GetItemMaster dynamoimpl
func (u *dynamoRepository) GetItemMaster(req Request) (Responce, error) {
	var items []itemMaster
	table := u.getTable()
	err := table.Get("user_id", req.UserID).Filter("group_id = ?", req.GroupID).All(&items) // Range("group_id", dynamo.Equal, req.GroupID).All(&items)
	if err != nil {
		u.logger.LogWrite(log.Error, "table.Get Error"+fmt.Sprint(err))
		return Responce{}, err
	}
	var list = make([]ItemMaster, 0)
	for _, item := range items {
		var t = NewItemMaster(item.UserID, item.GroupID, item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName)
		list = append(list, t)
	}
	return Responce{ItemMasters: list}, nil
}

// PutItemMaster dynamoimpl
func (u *dynamoRepository) PutItemMaster(req PutRequest) (PutResponce, error) {
	u.logger.LogWrite(log.Info, "start PutItemMaster")
	var batchSize = len(req.ItemMasters)
	var list = make([]interface{}, batchSize)
	for i := 0; i < batchSize; i++ {
		var item = req.ItemMasters[i]
		var t = NewitemMaster(item.UserID, item.GroupID, item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName)
		list[i] = t
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "batch insert", list)

	table := u.getTable()
	batch := table.Batch().Write().Put(list...)
	wrote, err := batch.Run()
	if err != nil {
		u.logger.LogWrite(log.Error, "error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutItemMaster")
		return PutResponce{Wrote: wrote}, err
	}
	if len(req.ItemMasters) != wrote {
		u.logger.LogWrite(log.Error, fmt.Sprintf("some wrote is failed. Total:%v Success:%v", len(req.ItemMasters), wrote))
	}
	u.logger.LogWrite(log.Info, "end PutItemMaster")
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
	table := db.Table("ItemObserveMaster")
	return table
}
