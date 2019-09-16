package controllers

// Request input
type Request struct {
	// ユーザーID
	UserID string `json:"user_id"`
}

// Responce output
type Responce struct {
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster `json:"item_masters"`
}

// ItemMaster リクエストの出力パラメータ
type ItemMaster struct {
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
func NewItemMaster(groupID string, productID string, storeType string, thretholdPrice int, itemName string) ItemMaster {
	return ItemMaster{
		GroupID:        groupID,
		ProductID:      productID,
		StoreType:      storeType,
		ThretholdPrice: thretholdPrice,
		ItemName:       itemName,
	}
}
