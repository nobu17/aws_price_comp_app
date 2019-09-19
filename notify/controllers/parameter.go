package controllers

// Request request common
type Request struct {
	// Method type
	Method string `json:"method"`
	// PutRequest
	PutRequest PutRequest `json:"put_param"`
}

// PutRequest struct
type PutRequest struct {
	// 送信あて先
	UserInfo UserInfo `json:"user_info"`
	// ItemGroup
	GroupID string `json:"group_id"`
	// 送信商品リスト
	ProductInfoList []ProductInfo `json:"product_info_list"`
}

// PutResponce result
type PutResponce struct {
}

// UserInfo information of send mail target.
type UserInfo struct {
	// UserID
	UserID string `json:"user_id"`
	// 名前
	Name string `json:"name"`
	// mail address
	Mail string `json:"mail"`
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
