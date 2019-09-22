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
	boffRepsotiroy    repositories.GetProductPriceImpl
}

// NewRepsoitoryFactrory constructor
func NewRepsoitoryFactrory(amaRepsotiroy, surugayaRepsitory, boffRepsotiroy repositories.GetProductPriceImpl) FactoryImpl {
	return &repsoitoryFactrory{
		amaRepsotiroy:     amaRepsotiroy,
		surugayaRepsitory: surugayaRepsitory,
		boffRepsotiroy:    boffRepsotiroy,
	}
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
	if info.IsStoreTypeBookoff(storeType) {
		return u.boffRepsotiroy, nil
	}
	return nil, errors.New("no match store type:" + storeType)
}
