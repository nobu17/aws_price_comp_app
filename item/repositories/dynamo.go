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
		var t = NewItemMaster(item.GroupID, item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName)
		list = append(list, t)
	}
	return Responce{ItemMasters: list}, nil
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
