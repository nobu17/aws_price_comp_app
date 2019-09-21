package services

// ServiceImpl service interface
type ServiceImpl interface {
	GetItemMasters(req InputModel) (OutputModel, error)
	PutItemMasters(req PutInputModel) (PutOutputModel, error)
}

// InputModel input
type InputModel struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
}

// OutputModel output
type OutputModel struct {
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster
}

// PutInputModel input
type PutInputModel struct {
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster
}

// PutOutputModel output
type PutOutputModel struct {
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
