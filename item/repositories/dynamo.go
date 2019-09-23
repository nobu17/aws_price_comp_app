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

type itemRecord struct {
	// ユーザID
	UserID string `dynamo:"user_id"`
	// グループID
	GroupID string `dynamo:"group_id"`
	// item list
	ItemList []itemMaster `dynamo:"item_list"`
}

// itemMaster リクエストの出力パラメータ
type itemMaster struct {
	// 商品ID
	ProductID string `dynamo:"product_id"`
	// 店鋪種類
	StoreType string `dynamo:"store_type"`
	// しきい値
	ThretholdPrice int `dynamo:"threthold_price"`
	// アイテム名
	ItemName string `dynamo:"item_name"`
}

// NewitemMaster construtor
func NewitemMaster(productID string, storeType string, thretholdPrice int, itemName string) itemMaster {
	return itemMaster{
		ProductID:      productID,
		StoreType:      storeType,
		ThretholdPrice: thretholdPrice,
		ItemName:       itemName,
	}
}

// GetItemMaster dynamoimpl
func (u *dynamoRepository) GetItemMaster(req Request) (Responce, error) {
	var userItem itemRecord
	table := u.getTable()
	err := table.Get("user_id", req.UserID).Range("group_id", dynamo.Equal, req.GroupID).One(&userItem)
	if err != nil {
		u.logger.LogWrite(log.Error, "table.Get Error"+fmt.Sprint(err))
		return Responce{}, err
	}
	var list = make([]ItemMaster, 0)
	for _, item := range userItem.ItemList {
		var t = NewItemMaster(item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName)
		list = append(list, t)
	}
	return Responce{UserID: userItem.UserID, GroupID: userItem.GroupID, ItemMasters: list}, nil
}

// PutItemMaster dynamoimpl
func (u *dynamoRepository) PutItemMaster(req PutRequest) (PutResponce, error) {
	u.logger.LogWrite(log.Info, "start PutItemMaster")

	var list = make([]itemMaster, 0)
	for _, item := range req.ItemMasters {
		var t = NewitemMaster(item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName)
		list = append(list, t)
	}
	var input = itemRecord{UserID: req.UserID, GroupID: req.GroupID, ItemList: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "put", input)

	table := u.getTable()
	err := table.Put(input).Run()

	if err != nil {
		u.logger.LogWrite(log.Error, "error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutItemMaster")
		return PutResponce{}, err
	}
	u.logger.LogWrite(log.Info, "end PutItemMaster")
	return PutResponce{}, nil
}

// DeleteItemMaster dynamo impl
func (u *dynamoRepository) DeleteItemMaster(req DeleteRequest) (DeleteResponce, error) {
	table := u.getTable()
	err := table.Delete("user_id", req.UserID).Range("group_id", req.GroupID).Run()
	if err != nil {
		u.logger.LogWrite(log.Error, "table.Del Error"+fmt.Sprint(err))
		return DeleteResponce{}, err
	}
	return DeleteResponce{}, nil
}

func (u *dynamoRepository) getTable() dynamo.Table {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String(region),
	})
	table := db.Table("ItemMaster")
	return table
}
