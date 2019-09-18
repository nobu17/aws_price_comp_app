package repositories

// Request リクエスト入力パラメータ
type Request struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
}

// Responce result
type Responce struct {
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster
}

// ItemMaster リクエストの出力パラメータ
type ItemMaster struct {
	// グループID
	GroupID string
	// 商品ID
	ProductID string
	// 店鋪種類
	StoreType string
	// しきい値
	ThretholdPrice int
	// アイテム名
	ItemName string
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

// ItemMasterImpl interface
type ItemMasterImpl interface {
	GetItemMaster(req Request) (Responce, error)
}
