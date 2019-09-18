package factories

import (
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
	switch storeType {
	case "amazon":
		return u.amaRepsotiroy, nil
	case "surugaya":
		return u.surugayaRepsitory, nil
	default:
		return nil, errors.New("no match store type:" + storeType)
	}
}
