package repositories

// UserInfoImpl interface of user repository
type UserInfoImpl interface {
	GetUserInfo(req GetRequest) (GetResponce, error)
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
