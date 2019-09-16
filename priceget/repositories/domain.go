package repositories

// Request リクエスト入力パラメータ
type Request struct {
	// 商品ID
	ProductID string
}

// ProductInfo リクエストの出力パラメータ
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

// NewProductInfo constructor
func NewProductInfo(productID string) ProductInfo {
	return ProductInfo{
		ProductID:   productID,
		Price:       -1,
		ShippingFee: 0,
		IsSoldOut:   false,
	}
}

// GetProductPriceImpl 価格取得インフェース
type GetProductPriceImpl interface {
	// 価格取得
	GetProductPrice(req Request) (ProductInfo, error)
}
