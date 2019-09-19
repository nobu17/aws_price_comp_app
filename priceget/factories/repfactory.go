package factories

import (
	"common/info"
	"errors"
	"priceget/repositories"
	"strings"
)

// repsoitoryFactrory struct
type repsoitoryFactrory struct {
	amaRepsotiroy     repositories.GetProductPriceImpl
	surugayaRepsitory repositories.GetProductPriceImpl
}

// NewRepsoitoryFactrory constructor
func NewRepsoitoryFactrory(amaRepsotiroy repositories.GetProductPriceImpl, surugayaRepsitory repositories.GetProductPriceImpl) FactoryImpl {
	return &repsoitoryFactrory{
		amaRepsotiroy:     amaRepsotiroy,
		surugayaRepsitory: surugayaRepsitory}
}

// GetPriceGetRepository implimation
func (u *repsoitoryFactrory) GetPriceGetRepository(storeType string) (repositories.GetProductPriceImpl, error) {
	if storeType == "" || strings.TrimSpace(storeType) == "" {
		return nil, errors.New("no storeType")
	}
	if info.IsStoreTypeAmazon(storeType) {
		return u.amaRepsotiroy, nil
	}
	if info.IsStoreTypeSurugaya(storeType) {
		return u.surugayaRepsitory, nil
	}
	return nil, errors.New("no match store type:" + storeType)
}
