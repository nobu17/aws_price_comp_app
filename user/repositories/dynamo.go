package repositories

import (
	"common/log"
	"errors"
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
func NewDynamoRepository(logger log.LoggerImpl) UserInfoImpl {
	return &dynamoRepository{logger: logger}
}

// UserInfo user information
type userInfo struct {
	// ユーザーID
	UserID string `dynamo:"user_id"`
	// 名称
	Name string `dynamo:"name"`
	// メール
	Mail string `dynamo:"mail"`
}

// ItemGroup アイテムマスタのグループ情報
type itemGroup struct {
	// id
	GroupID string `dynamo:"group_id"`
	// name
	GroupName string `dynamo:"group_name"`
}

// GetAlertLog impl
func (u *dynamoRepository) GetUserInfo(req GetRequest) (GetResponce, error) {
	// get user
	var users []userInfo
	table := u.getUserTable()
	err := table.Get("user_id", req.UserID).Filter("password = ?", req.Password).All(&users)
	if err != nil {
		u.logger.LogWrite(log.Error, "usertable.Get Error"+fmt.Sprint(err))
		return GetResponce{}, err
	}
	// no user or auth error
	if len(users) < 1 {
		u.logger.LogWrite(log.Error, "no user or password wrong")
		return GetResponce{}, errors.New("no user or password wrong")
	}
	// get itemgroup
	var itemGroups []itemGroup
	table = u.getItemGroup()
	err = table.Get("user_id", req.UserID).All(&itemGroups)
	if err != nil {
		u.logger.LogWrite(log.Error, "itemGroupTable.Get Error"+fmt.Sprint(err))
		return GetResponce{}, err
	}
	// no group
	if len(itemGroups) < 1 {
		u.logger.LogWrite(log.Warn, "no itemGroups")
	}

	var list = make([]ItemGroup, 0)
	for _, item := range itemGroups {
		var t = NewItemGroup(item.GroupID, item.GroupName)
		list = append(list, t)
	}

	return GetResponce{UserInfo: NewUserInfo(users[0].UserID, users[0].Name, users[0].Mail), ItemGroupList: list}, nil
}

// DeleteItemGroup impl
func (u *dynamoRepository) DeleteItemGroup(req DeleteItemGroupRequest) (DeleteItemGroupResponce, error) {
	table := u.getItemGroup()
	var successGroups = make([]string, 0)
	var faileGroups = make([]string, 0)
	for _, group := range req.GroupIDList {
		err := table.Delete("user_id", req.UserID).Range("group_id", group).Run()
		if err != nil {
			u.logger.LogWrite(log.Error, "table.Del Error"+fmt.Sprint(err))
			faileGroups = append(faileGroups, group)
			continue
		}
		successGroups = append(successGroups, group)
	}
	return DeleteItemGroupResponce{SuccessItemGroupList: successGroups, FailedItemGroupList: faileGroups}, nil
}

func (u *dynamoRepository) getUserTable() dynamo.Table {
	db := u.getDB()
	table := db.Table("UserMaster")
	return table
}

func (u *dynamoRepository) getItemGroup() dynamo.Table {
	db := u.getDB()
	table := db.Table("ItemGroupMaster")
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
