package controllers

import (
	"common/log"
	"notify/repositories"
	"notify/services"
)

// Container struct
type Container struct {
}

// NewService init
func (u *Container) NewService() services.ServiceImpl {
	logger := u.NewLogger()
	repo := repositories.NewSesRepository(logger)
	serv := services.NewItemMasterService(logger, repo)

	return serv
}

// NewLogger make a new looger
func (u *Container) NewLogger() log.LoggerImpl {
	logger := log.NewAwsLogger()
	return logger
}
