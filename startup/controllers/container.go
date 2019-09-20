package controllers

import (
	"common/log"
	"startup/repositories"
	"startup/services"
)

// Container struct
type Container struct {
}

// NewService init
func (u *Container) NewService() services.ServiceImpl {
	logger := u.NewLogger()
	lRepo := repositories.NewLambdaRepositories(logger)
	sRepo := repositories.NewSQSRepositories(logger)
	serv := services.NewObserveService(logger, lRepo, sRepo)

	return serv
}

// NewLogger make a new looger
func (u *Container) NewLogger() log.LoggerImpl {
	logger := log.NewAwsLogger()
	return logger
}
