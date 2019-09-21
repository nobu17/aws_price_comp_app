package repositories

// NotifyImpl item master get impl
type NotifyImpl interface {
	SendNotify(req NotifyPutRequest) (NotifyPutResponce, error)
}

// NotifyPutRequest struct
type NotifyPutRequest struct {
	// 送信あて先
	UserInfo NotifyUserInfo `json:"user_info"`
	// ItemGroup
	GroupID string `json:"group_id"`
	// 送信商品リスト
	ProductInfoList []NotifyProductInfo `json:"product_info_list"`
}

// NotifyPutResponce result
type NotifyPutResponce struct {
}

// NotifyUserInfo information of send mail target.
type NotifyUserInfo struct {
	// UserID
	UserID string `json:"user_id"`
	// 名前
	Name string `json:"name"`
	// mail address
	Mail string `json:"mail"`
}

// NewNotifyUserInfo constructor.
func NewNotifyUserInfo(userID string, name string, mail string) NotifyUserInfo {
	return NotifyUserInfo{
		UserID: userID,
		Name:   name,
		Mail:   mail,
	}
}

// NotifyProductInfo struct
type NotifyProductInfo struct {
	// 商品ID
	ProductID string `json:"product_id"`
	// 店鋪種類
	StoreType string `json:"store_type"`
	// 商品名
	Name string `json:"name"`
	// 価格
	Price int `json:"price"`
	// 送料
	ShippingFee int `json:"shipping_fee"`
}

// NewNotifyProductInfo construcotr
func NewNotifyProductInfo(productID string, storeType string, name string, price int, shippingFee int) NotifyProductInfo {
	return NotifyProductInfo{
		ProductID:   productID,
		StoreType:   storeType,
		Name:        name,
		Price:       price,
		ShippingFee: shippingFee,
	}
}
