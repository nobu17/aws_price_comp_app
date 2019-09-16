package controllers

import (
	"common/log"
	"errors"
	"fmt"
	"item/alert/services"
	"strings"
	"time"
)

// AlertController controller
type AlertController struct {
	logger  log.LoggerImpl
	service services.ServiceImpl
}

// NewAlertController constroctor
func NewAlertController() AlertController {
	container := Container{}
	ser := container.NewService()
	log := container.NewLogger()
	return AlertController{
		logger:  log,
		service: ser,
	}
}

// GetAlertLog getitem
func (u *AlertController) GetAlertLog(req GetRequest) (GetResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start GetAlertLog", req)

	err := u.validate(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end GetAlertLog:")
		return GetResponce{}, err
	}

	var inputModel = services.GetInputModel{UserID: req.UserID, MinAlertDate: req.MinAlertDate}
	res, err := u.service.GetAlertLog(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end GetAlertLog:")
		return GetResponce{}, err
	}
	var list = make([]SendAlertLog, 0)
	for _, item := range res.SendAlertLogList {
		list = append(list, NewSendAlertLog(item.UserID, item.AlertDate, item.StoreType, item.ProductID, item.Price))
	}

	var respo = GetResponce{SendAlertLogList: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetAlertLog:", respo)
	return respo, nil
}

func (u *AlertController) validate(req GetRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	_, err := time.Parse("20060102", req.MinAlertDate)
	if err != nil {
		return errors.New("MinAlertDate is format error:" + req.MinAlertDate)
	}

	return nil
}
