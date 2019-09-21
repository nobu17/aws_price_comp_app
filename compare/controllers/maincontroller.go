package controllers

import (
	"common/log"
	"compare/services"
	"errors"
	"fmt"
	"strings"
)

// CompareController controller
type CompareController struct {
	logger  log.LoggerImpl
	service services.ServiceImpl
}

// NewCompareController constroctor.
func NewCompareController() CompareController {
	container := Container{}
	ser := container.NewService()
	log := container.NewLogger()
	return CompareController{
		logger:  log,
		service: ser,
	}
}

// StartCompare get user
func (u *CompareController) StartCompare(req Request) (Responce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start CompareController:StartCompare", req)

	err := u.validateReq(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end StartupController:StartObserv")
		return Responce{}, err
	}

	var inputModel = services.NewInputModel(req.UserID, req.UserName, req.GroupID, req.GroupName, req.Mail)
	_, err = u.service.StartCompare(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end StartupController:StartCompare")
		return Responce{}, err
	}
	return Responce{}, nil
}

func (u *CompareController) validateReq(req Request) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.UserName == "" || strings.TrimSpace(req.UserName) == "" {
		return errors.New("UserName is empty")
	}
	if req.Mail == "" || strings.TrimSpace(req.Mail) == "" {
		return errors.New("Mail is empty")
	}
	if req.GroupID == "" || strings.TrimSpace(req.GroupID) == "" {
		return errors.New("GroupID is empty")
	}
	if req.GroupName == "" || strings.TrimSpace(req.GroupName) == "" {
		return errors.New("GroupName is empty")
	}
	return nil
}
