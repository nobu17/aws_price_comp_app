package controllers

import (
	"common/log"
	"user/repositories"
	"user/services"
)

// Container struct
type Container struct {
}

// NewService init
func (u *Container) NewService() services.ServiceImpl {
	logger := u.NewLogger()
	repo := repositories.NewDynamoRepository(logger)
	serv := services.NewUserService(logger, repo)

	return serv
}

// NewLogger make a new looger
func (u *Container) NewLogger() log.LoggerImpl {
	logger := log.NewAwsLogger()
	return logger
}
