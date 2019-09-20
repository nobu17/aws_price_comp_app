package repositories

// SendObservImpl interface of user repository
type SendObservImpl interface {
	SendObservRequest(req SendRequest) error
}

// SendRequest send que resuest.
type SendRequest struct {
	// ユーザー情報
	UserInfo UserInfo `json:"user_info"`
	// アイテムグループのマスタ
	ItemGroupList []ItemGroup `json:"item_group_list"`
}