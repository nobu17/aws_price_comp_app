package repositories

// ItemImpl item master get impl
type ItemImpl interface {
	GetItems(req ItemGetRequest) (ItemGetResponce, error)
}

// ItemGetRequest struct for item getting
type ItemGetRequest struct {
	// userid
	UserID string `json:"user_id"`
	// id
	GroupID string `json:"group_id"`
}

// ItemGetResponce output
type ItemGetResponce struct {
	// userid
	UserID string `json:"user_id"`
	// id
	GroupID string `json:"group_id"`
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster `json:"item_masters"`
}

// ItemMaster リクエストの出力パラメータ
type ItemMaster struct {
	// 商品ID
	ProductID string `json:"product_id"`
	// 店鋪種類
	StoreType string `json:"store_type"`
	// しきい値
	ThretholdPrice int `json:"threthold_price"`
	// アイテム名
	ItemName string `json:"item_name"`
}
