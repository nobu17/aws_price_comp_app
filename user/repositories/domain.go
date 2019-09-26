package repositories

// UserInfoImpl interface of user repository
type UserInfoImpl interface {
	GetUserInfo(req GetRequest) (GetResponce, error)
	PutItemGroup(req PutItemGroupRequest) (PutItemGroupResponce, error)
	DeleteItemGroup(req DeleteItemGroupRequest) (DeleteItemGroupResponce, error)
}
// GroupInfoImpl interface of user repository
type GroupInfoImpl interface {
	DeleteItemGroup(req DeleteItemGroupRequest) (DeleteItemGroupResponce, error)
}

// GetRequest リクエスト入力パラメータ
type GetRequest struct {
	// ユーザーID
	UserID string
	// パスワード
	Password string
}

// NewGetRequest constructor
func NewGetRequest(userID string, password string) GetRequest {
	return GetRequest{UserID: userID, Password: password}
}

// GetResponce 応答
type GetResponce struct {
	// ユーザー情報
	UserInfo UserInfo
	// アイテムグループのマスタ
	ItemGroupList []ItemGroup
}

// UserInfo user information
type UserInfo struct {
	// ユーザーID
	UserID string
	// 名称
	Name string
	// メール
	Mail string
}

// NewUserInfo constructor
func NewUserInfo(userID string, name string, mail string) UserInfo {
	return UserInfo{UserID: userID, Name: name, Mail: mail}
}

// PutItemGroupRequest リクエスト入力パラメータ
type PutItemGroupRequest struct {
	// ユーザーID
	UserID string
	// GroupIDs
	GroupList []ItemGroup
}

// PutItemGroupResponce リクエスト出力パラメータ
type PutItemGroupResponce struct {
	// 削除失敗したアイテムグループのマスタ
	SuccessItemGroupList []string
	// 削除失敗したアイテムグループのマスタ
	FailedPutGroupList []string
}

// DeleteItemGroupRequest リクエスト入力パラメータ
type DeleteItemGroupRequest struct {
	// ユーザーID
	UserID string
	// GroupIDs
	GroupIDList []string
}

// NewDeleteItemGroupRequest constructor
func NewDeleteItemGroupRequest(userID string, groupIDList []string) DeleteItemGroupRequest {
	return DeleteItemGroupRequest{UserID: userID, GroupIDList: groupIDList}
}

// DeleteItemGroupResponce リクエスト出力パラメータ
type DeleteItemGroupResponce struct {
	// 削除失敗したアイテムグループのマスタ
	SuccessItemGroupList []string
	// 削除失敗したアイテムグループのマスタ
	FailedItemGroupList []string
}

// ItemGroup アイテムマスタのグループ情報
type ItemGroup struct {
	// id
	GroupID string
	// name
	GroupName string
}

// NewItemGroup constructor
func NewItemGroup(groupID string, groupName string) ItemGroup {
	return ItemGroup{GroupID: groupID, GroupName: groupName}
}
