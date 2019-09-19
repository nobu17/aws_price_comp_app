package services

// ServiceImpl service interface
type ServiceImpl interface {
	SendNotify(req PutInputModel) (PutOutputModel, error)
}

// PutInputModel struct
type PutInputModel struct {
	// 送信あて先
	UserInfo UserInfo
	// ItemGroup
	GroupID string
	// 送信商品リスト
	ProductInfoList []ProductInfo
}

// PutOutputModel struct
type PutOutputModel struct {
}

// UserInfo information of send mail target.
type UserInfo struct {
	// UserID
	UserID string
	// 名前
	Name string
	// mail address
	Mail string
}

// NewUserInfo constructor.
func NewUserInfo(userID string, name string, mail string) UserInfo {
	return UserInfo{
		UserID: userID,
		Name:   name,
		Mail:   mail,
	}
}

// ProductInfo struct
type ProductInfo struct {
	// 商品ID
	ProductID string
	// 店鋪種類
	StoreType string
	// 商品名
	Name string
	// 価格
	Price int
	// 送料
	ShippingFee int
}

// NewProductInfo construcotr
func NewProductInfo(productID string, storeType string, name string, price int, shippingFee int) ProductInfo {
	return ProductInfo{
		ProductID:   productID,
		StoreType:   storeType,
		Name:        name,
		Price:       price,
		ShippingFee: shippingFee,
	}
}
