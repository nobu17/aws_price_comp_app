package controllers

// Request struct
type Request struct {
	// 商品リスト
	ProductList []ProductRequest `json:"productList"`
}

// ProductRequest request
type ProductRequest struct {
	// 店鋪種類
	StoreType string `json:"storeType"`
	// 商品ID
	ProductID string `json:"productID"`
}

// Responce struct
type Responce struct {
	// 商品リスト
	ProductInfoList []ProductInfo `json:"ProductInfoList"`
	// 失敗リスト
	FailProductInfoList []ProductInfo `json:"FailProductInfoList"`
}

// ProductInfo product information
type ProductInfo struct {
	// 商品ID
	ProductID string `json:"productID"`
	// 価格
	Price int `json:"price"`
	// 送料
	ShippingFee int `json:"shippingFee"`
	// 売り切れ
	IsSoldOut bool `json:"isSoldOut"`
}
