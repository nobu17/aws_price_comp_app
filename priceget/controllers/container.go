package controllers

import (
	"common/log"
	"priceget/factories"
	"priceget/repositories"
	"priceget/services"
)

// Container struct
type Container struct {
}

// NewService init
func (u *Container) NewService() services.ServiceImpl {
	logger := u.NewLogger()
	aRepo := repositories.NewAmazonMakertRepository(logger)
	sRepo := repositories.NewSurugayaRepository(logger)
	factory := factories.NewRepsoitoryFactrory(aRepo, sRepo)
	serv := services.NewProductService(factory, logger)

	return serv
}

// NewLogger make a new looger
func (u *Container) NewLogger() log.LoggerImpl {
	logger := log.NewAwsLogger()
	return logger
}
