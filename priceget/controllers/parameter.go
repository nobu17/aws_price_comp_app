package controllers

// Request struct
type Request struct {
	// 商品リスト
	ProductList []ProductRequest `json:"product_list"`
}

// ProductRequest request
type ProductRequest struct {
	// 店鋪種類
	StoreType string `json:"store_type"`
	// 商品ID
	ProductID string `json:"product_id"`
}

// Responce struct
type Responce struct {
	// 商品リスト
	ProductInfoList []ProductInfo `json:"product_infolist"`
	// 失敗リスト
	FailProductInfoList []ProductInfo `json:"fail_product_infolist"`
}

// ProductInfo product information
type ProductInfo struct {
	// 商品ID
	ProductID string `json:"product_id"`
	// 商品type
	StoreType string `json:"store_type"`
	// 価格
	Price int `json:"price"`
	// 送料
	ShippingFee int `json:"shipping_fee"`
	// 売り切れ
	IsSoldOut bool `json:"is_soldout"`
}
