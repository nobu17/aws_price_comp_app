package services

// ServiceImpl service interface
type ServiceImpl interface {
	GetItemMasters(req InputModel) (OutputModel, error)
	PutItemMasters(req PutInputModel) (PutOutputModel, error)
	DeleteItemMasters(req DeleteInputModel) (DeleteOutputModel, error)
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
	// ユーザーID
	UserID string
	// グループID
	GroupID string
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster
}

// PutInputModel input
type PutInputModel struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
	// ItemMasters 商品リスト
	ItemMasters []ItemMaster
}

// PutOutputModel output
type PutOutputModel struct {
}

// DeleteInputModel input
type DeleteInputModel struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
}

// DeleteOutputModel output
type DeleteOutputModel struct {
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
