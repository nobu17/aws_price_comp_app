package controllers

// Request requestcommon
type Request struct {
	// method param
	Method string `json:"method"`
	// GetRequest method get case getparam
	GetRequest GetRequest `json:"get_param"`
	// PutRequest method put case getparam
	PutRequest PutRequest `json:"put_param"`
	// DeleteRequest method put case delparam
	DeleteRequest DeleteRequest `json:"delete_param"`
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
	// ユーザーID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster `json:"item_masters"`
}

// PutRequest input
type PutRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster `json:"item_masters"`
}

// PutResponce output
type PutResponce struct {
}

// DeleteRequest input
type DeleteRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
}

// DeleteResponce output
type DeleteResponce struct {
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

// NewItemMaster construtor
func NewItemMaster(productID string, storeType string, thretholdPrice int, itemName string) ItemMaster {
	return ItemMaster{
		ProductID:      productID,
		StoreType:      storeType,
		ThretholdPrice: thretholdPrice,
		ItemName:       itemName,
	}
}
