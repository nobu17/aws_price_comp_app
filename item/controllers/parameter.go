package controllers

// Request requestcommon
type Request struct {
	// method param
	Method string `json:"method"`
	// GetRequest method get case getparam
	GetRequest GetRequest `json:"get_param"`
	// PutRequest method put case getparam
	PutRequest PutRequest `json:"put_param"`
}

// GetRequest input
type GetRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
}

// GetResponce output
type GetResponce struct {
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster `json:"item_masters"`
}

// PutRequest input
type PutRequest struct {
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster `json:"item_masters"`
}

// PutResponce output
type PutResponce struct {
}

// ItemMaster リクエストの出力パラメータ
type ItemMaster struct {
	// ユーザID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
	// 商品ID
	ProductID string `json:"product_id"`
	// 店鋪種類
	StoreType string `json:"store_type"`
	// しきい値
	ThretholdPrice int `json:"threthold_price"`
	// アイテム名
	ItemName string `json:"item_name"`
}

// NewItemMaster construtor
func NewItemMaster(userID string, groupID string, productID string, storeType string, thretholdPrice int, itemName string) ItemMaster {
	return ItemMaster{
		UserID:         userID,
		GroupID:        groupID,
		ProductID:      productID,
		StoreType:      storeType,
		ThretholdPrice: thretholdPrice,
		ItemName:       itemName,
	}
}
