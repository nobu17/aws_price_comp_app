package controllers

import (
	"common/log"
	"compare/repositories"
	"compare/services"
)

// Container struct
type Container struct {
}

// NewService init
func (u *Container) NewService() services.ServiceImpl {
	logger := u.NewLogger()
	aRepo := repositories.NewAlertambdaRepositories(logger)
	iRepo := repositories.NewItemLambdaRepositories(logger)
	nRepo := repositories.NewNotifyLambdaRepositories(logger)
	pRepo := repositories.NewPriceLambdaRepositories(logger)
	serv := services.NewCompareService(logger, aRepo, iRepo, pRepo, nRepo)

	return serv
}

// NewLogger make a new looger
func (u *Container) NewLogger() log.LoggerImpl {
	logger := log.NewAwsLogger()
	return logger
}
