package controllers

// Request request common
type Request struct {
	// Method type
	Method string `json:"method"`
	// GetParam
	GetRequest GetRequest `json:"get_param"`
	// DeleteParam
	DeleteGroupRequest DeleteGroupRequest `json:"delete_group_param"`
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

// DeleteGroupRequest リクエスト入力パラメータ
type DeleteGroupRequest struct {
	// ユーザーID
	UserID string
	// GroupIDs
	GroupIDList []string
}

// DeleteGroupResponce リクエスト出力パラメータ
type DeleteGroupResponce struct {
	// 削除失敗したアイテムグループのマスタ
	SuccessItemGroupList []string
	// 削除失敗したアイテムグループのマスタ
	FailedDeleteGroupList []string
	// 削除失敗したアイテムのマスタ
	FailedDeleteItemList []string
}

// NewDeleteGroupResponce constructor
func NewDeleteGroupResponce(successItemGroupList, failedDeleteGroupList, failedDeleteItemList []string) DeleteGroupResponce {
	return DeleteGroupResponce{SuccessItemGroupList: successItemGroupList, FailedDeleteGroupList: failedDeleteGroupList, FailedDeleteItemList: failedDeleteItemList}
}
