package services

// ServiceImpl service interface
type ServiceImpl interface {
	GetUserInfo(req GetInputModel) (GetOutputModel, error)
	PutItemGroup(req PutItemGroupInputModel) (PutItemGroupOutputModel, error) 
	DeleteItemGroup(req DeleteInputModel) (DeleteOutputModel, error)
}

// GetInputModel input
type GetInputModel struct {
	// ユーザーID
	UserID string
	// パスワード
	Password string
}

// NewGetInputModel constructor
func NewGetInputModel(userID string, password string) GetInputModel {
	return GetInputModel{UserID: userID, Password: password}
}

// GetOutputModel output
type GetOutputModel struct {
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

// PutItemGroupInputModel リクエスト入力パラメータ
type PutItemGroupInputModel struct {
	// ユーザーID
	UserID string
	// GroupIDs
	GroupList []ItemGroup
}

// PutItemGroupOutputModel リクエスト出力パラメータ
type PutItemGroupOutputModel struct {
	// 削除失敗したアイテムグループのマスタ
	SuccessItemGroupList []string
	// 削除失敗したアイテムグループのマスタ
	FailedPutGroupList []string
}

// DeleteInputModel リクエスト入力パラメータ
type DeleteInputModel struct {
	// ユーザーID
	UserID string
	// GroupIDs
	GroupIDList []string
}

// NewDeleteInputModel constructor
func NewDeleteInputModel(userID string, groupIDList []string) DeleteInputModel {
	return DeleteInputModel{UserID: userID, GroupIDList: groupIDList}
}

// DeleteOutputModel リクエスト出力パラメータ
type DeleteOutputModel struct {
	// 削除失敗したアイテムグループのマスタ
	SuccessItemGroupList []string
	// 削除失敗したアイテムグループのマスタ
	FailedDeleteGroupList []string
	// 削除失敗したアイテムのマスタ
	FailedDeleteItemList []string
}

// NewDeleteOutputModel constructor
func NewDeleteOutputModel(successItemGroupList, failedDeleteGroupList, failedDeleteItemList []string) DeleteOutputModel {
	return DeleteOutputModel{SuccessItemGroupList: successItemGroupList, FailedDeleteGroupList: failedDeleteGroupList, FailedDeleteItemList: failedDeleteItemList}
}
