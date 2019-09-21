package repositories

// ItemMasterImpl interface
type ItemMasterImpl interface {
	GetItemMaster(req Request) (Responce, error)
	PutItemMaster(req PutRequest) (PutResponce, error)
}

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

// PutRequest リクエスト入力パラメータ
type PutRequest struct {
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster
}

// PutResponce result
type PutResponce struct {
	// 書き込み数
	Wrote int
}

// ItemMaster リクエストの出力パラメータ
type ItemMaster struct {
	// ユーザID
	UserID string
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
