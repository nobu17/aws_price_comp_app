package factories

import "priceget/repositories"

// FactoryImpl interface for repository
type FactoryImpl interface {
	GetPriceGetRepository(storeType string) (repositories.GetProductPriceImpl, error)
}
