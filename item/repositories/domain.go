package repositories

// ItemMasterImpl interface
type ItemMasterImpl interface {
	GetItemMaster(req Request) (Responce, error)
	PutItemMaster(req PutRequest) (PutResponce, error)
	DeleteItemMaster(req DeleteRequest) (DeleteResponce, error)
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
	// ユーザーID
	UserID string
	// グループID
	GroupID string
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster
}

// PutRequest リクエスト入力パラメータ
type PutRequest struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster
}

// PutResponce result
type PutResponce struct {
}

// DeleteRequest delete req.
type DeleteRequest struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
}

// DeleteResponce result
type DeleteResponce struct {
}

// ItemMaster リクエストの出力パラメータ
type ItemMaster struct {
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
func NewItemMaster(productID string, storeType string, thretholdPrice int, itemName string) ItemMaster {
	return ItemMaster{
		ProductID:      productID,
		StoreType:      storeType,
		ThretholdPrice: thretholdPrice,
		ItemName:       itemName,
	}
}
