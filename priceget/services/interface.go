package services

import (
	"common/log"
	"priceget/factories"
)

// ServiceImpl impl
type ServiceImpl interface {
	SetImpl(factory factories.FactoryImpl, logger log.LoggerImpl) error
	GetProductInfo(req InputProductModel) (OutputProductModel, error)
}

// InputProductModel 入力モデル
type InputProductModel struct {
	// 商品リスト
	ProductList []ProductRequest
}

// ProductRequest request
type ProductRequest struct {
	// 店鋪種類
	StoreType string
	// 商品ID
	ProductID string
}

// OutputProductModel 出力モデル
type OutputProductModel struct {
	// 商品リスト
	ProductInfoList []ProductInfo
	// 失敗リスト
	FailProductInfoList []ProductInfo
}

// ProductInfo product information
type ProductInfo struct {
	// 商品ID
	ProductID string
	// 価格
	Price int
	// 送料
	ShippingFee int
	// 売り切れ
	IsSoldOut bool
}
