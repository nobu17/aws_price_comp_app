package controllers

import (
	"common/log"
	"errors"
	"fmt"
	"startup/services"
	"strings"
)

// StartupController controller
type StartupController struct {
	logger  log.LoggerImpl
	service services.ServiceImpl
}

// NewStartupController constroctor
func NewStartupController() StartupController {
	container := Container{}
	ser := container.NewService()
	log := container.NewLogger()
	return StartupController{
		logger:  log,
		service: ser,
	}
}

// StartObserv get user
func (u *StartupController) StartObserv(req Request) (Responce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start StartupController:StartObserv", req)

	err := u.validateReq(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end StartupController:StartObserv")
		return Responce{}, err
	}

	var inputModel = services.NewInputModel(req.UserID, req.Password)
	_, err = u.service.StartObserve(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end UserController:GetUserInfo")
		return Responce{}, err
	}
	return Responce{}, nil
}

func (u *StartupController) validateReq(req Request) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.Password == "" || strings.TrimSpace(req.Password) == "" {
		return errors.New("Password is empty")
	}
	return nil
}
