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
	err := table.Get("user_id", req.UserID).All(&items)
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
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	table := db.Table("ItemMaster")
	return table
}
