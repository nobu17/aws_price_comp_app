package controllers

// Request request common
type Request struct {
	// Method type
	Method string `json:"method"`
	// GetParam
	GetRequest GetRequest `json:"get_param"`
}

// GetRequest リクエスト入力パラメータ
type GetRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// パスワード
	Password string `json:"password"`
}

// NewGetRequest constructor
func NewGetRequest(userID string, password string) GetRequest {
	return GetRequest{UserID: userID, Password: password}
}

// GetResponce 応答
type GetResponce struct {
	// ユーザー情報
	UserInfo UserInfo `json:"user_info"`
	// アイテムグループのマスタ
	ItemGroupList []ItemGroup `json:"item_group_list"`
}

// UserInfo user information
type UserInfo struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// 名称
	Name string `json:"name"`
	// メール
	Mail string `json:"mail"`
}

// NewUserInfo constructor
func NewUserInfo(userID string, name string, mail string) UserInfo {
	return UserInfo{UserID: userID, Name: name, Mail: mail}
}

// ItemGroup アイテムマスタのグループ情報
type ItemGroup struct {
	// id
	GroupID string `json:"group_id"`
	// name
	GroupName string `json:"group_name"`
}

// NewItemGroup constructor
func NewItemGroup(groupID string, groupName string) ItemGroup {
	return ItemGroup{GroupID: groupID, GroupName: groupName}
}
