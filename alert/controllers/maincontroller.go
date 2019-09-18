package controllers

import (
	"alert/services"
	"common/log"
	"errors"
	"fmt"
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
	u.logger.LogWriteWithMsgAndObj(log.Info, "start AlertController:GetAlertLog", req)

	err := u.validateGetReq(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end AlertController:GetAlertLog:")
		return GetResponce{}, err
	}

	var inputModel = services.GetInputModel{UserID: req.UserID, MinAlertDate: req.MinAlertDate}
	res, err := u.service.GetAlertLog(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end AlertController:GetAlertLog:")
		return GetResponce{}, err
	}
	var list = make([]SendAlertLog, 0)
	for _, item := range res.SendAlertLogList {
		list = append(list, NewSendAlertLog(item.UserID, item.AlertDate, item.StoreType, item.ProductID, item.Price))
	}

	var respo = GetResponce{SendAlertLogList: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end AlertController:GetAlertLog:", respo)
	return respo, nil
}

func (u *AlertController) validateGetReq(req GetRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	_, err := time.Parse("20060102", req.MinAlertDate)
	if err != nil {
		return errors.New("MinAlertDate is format error:" + req.MinAlertDate)
	}

	return nil
}

// PutAlertLog put alertlist
func (u *AlertController) PutAlertLog(req PutRequest) (PutResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start PutAlertLog", req)

	err := u.validatePutReq(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutAlertLog:")
		return PutResponce{}, err
	}

	var list = make([]services.SendAlertLog, 0)
	for _, item := range req.PutAlertLogList {
		list = append(list, services.NewSendAlertLog(item.UserID, item.AlertDate, item.StoreType, item.ProductID, item.Price))
	}
	var inputModel = services.PutInputModel{PutAlertLogList: list}
	err = u.service.PutAlertLog(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutAlertLog:")
		return PutResponce{}, err
	}

	var respo = PutResponce{}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end PutAlertLog:", respo)
	return PutResponce{}, nil
}

func (u *AlertController) validatePutReq(req PutRequest) error {
	if len(req.PutAlertLogList) == 0 {
		return errors.New("PutAlertLogList is empty")
	}

	return nil
}
