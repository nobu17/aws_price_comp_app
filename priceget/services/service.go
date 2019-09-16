package services

import (
	"common/log"
	"errors"
	"priceget/factories"
	"priceget/repositories"
)

// productService serivce
type productService struct {
	logger     log.LoggerImpl
	factory    factories.FactoryImpl
	repository repositories.GetProductPriceImpl
}

// NewProductService constructor
func NewProductService(factory factories.FactoryImpl, logger log.LoggerImpl) ServiceImpl {
	instance := productService{}
	instance.SetImpl(factory, logger)
	return &instance
}

// SetImpl setinterfaces
func (u *productService) SetImpl(factory factories.FactoryImpl, logger log.LoggerImpl) error {
	if factory == nil {
		return errors.New("factory is null")
	}
	if logger == nil {
		return errors.New("logger is null")
	}
	u.factory = factory
	u.logger = logger
	return nil
}

// GetProductInfo service impl
func (u *productService) GetProductInfo(req InputProductModel) (OutputProductModel, error) {
	if req.ProductList == nil || len(req.ProductList) == 0 {
		return OutputProductModel{}, errors.New("ProductList is empty")
	}
	var list = make([]ProductInfo, 0)
	var failList = make([]ProductInfo, 0)
	for _, prod := range req.ProductList {
		repository, err := u.factory.GetPriceGetRepository(prod.StoreType)
		if err != nil {
			return OutputProductModel{}, errors.New("repository get error")
		}
		u.repository = repository
		res, err := u.repository.GetProductPrice(repositories.Request{ProductID: prod.ProductID})
		if err != nil {
			failList = append(failList, ProductInfo{ProductID: res.ProductID, Price: res.Price, ShippingFee: res.ShippingFee, IsSoldOut: res.IsSoldOut})
		} else {
			list = append(list, ProductInfo{ProductID: res.ProductID, Price: res.Price, ShippingFee: res.ShippingFee, IsSoldOut: res.IsSoldOut})
		}
	}
	return OutputProductModel{ProductInfoList: list, FailProductInfoList: failList}, nil
}
