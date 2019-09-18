package services

// ServiceImpl service interface
type ServiceImpl interface {
	GetUserInfo(req GetInputModel) (GetOutputModel, error)
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
