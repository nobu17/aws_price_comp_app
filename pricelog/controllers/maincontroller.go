package controllers

import (
	"common/log"
	"errors"
	"fmt"
	"pricelog/services"
	"strings"
)

// PriceLogController controller
type PriceLogController struct {
	logger  log.LoggerImpl
	service services.ServiceImpl
}

// NewPriceLogController constroctor
func NewPriceLogController() PriceLogController {
	container := Container{}
	ser := container.NewService()
	log := container.NewLogger()
	return PriceLogController{
		logger:  log,
		service: ser,
	}
}

// GetPriceLog getpricelogs
func (u *PriceLogController) GetPriceLog(req GetRequest) (GetResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start GetPriceLog", req)

	err := u.validateGetReq(req)
	if err != nil {
		u.logger.LogWrite(log.Error, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end GetPriceLog:")
		return GetResponce{}, err
	}

	var inputModel = services.GetInputModel{UserID: req.UserID, GroupID: req.GroupID}
	res, err := u.service.GetPriceLog(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Error, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end GetPriceLog:")
		return GetResponce{}, err
	}
	var list = make([]PriceLogForGet, 0)
	for _, item := range res.PriceLogList {
		list = append(list, PriceLogForGet{StoreType: item.StoreType, ItemID: item.ItemID, Price: item.Price, LastModifiedDatetime: item.LastModifiedDatetime})
	}

	var respo = GetResponce{UserID: req.UserID, GroupID: req.GroupID, PriceLogList: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetPriceLog:", respo)
	return respo, nil
}

// PutPriceLog putpircelosg
func (u *PriceLogController) PutPriceLog(req PutRequest) (PutResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start PutPriceLog", req)

	err := u.validatePutReq(req)
	if err != nil {
		u.logger.LogWrite(log.Error, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end PutPriceLog:")
		return PutResponce{}, err
	}
	list := make([]services.PriceLog, 0)
	for _, item := range req.PriceLogList {
		list = append(list, services.PriceLog{StoreType: item.StoreType, ItemID: item.ItemID, Price: item.Price})
	}

	if isPutMode(req.Mode) {
		var inputModel = services.PutInputModel{UserID: req.UserID, GroupID: req.GroupID, PriceLogList: list}
		_, err := u.service.PutPriceLog(inputModel)
		if err != nil {
			u.logger.LogWrite(log.Error, "servie error:"+fmt.Sprint(err))
			u.logger.LogWrite(log.Error, "end PutPriceLog:")
			return PutResponce{}, err
		}
	} else if isUpdateMode(req.Mode) {
		var inputModel = services.UpdateInputModel{UserID: req.UserID, GroupID: req.GroupID, PriceLogList: list}
		_, err := u.service.UpdatePriceLog(inputModel)
		if err != nil {
			u.logger.LogWrite(log.Error, "servie error:"+fmt.Sprint(err))
			u.logger.LogWrite(log.Error, "end PutPriceLog:")
			return PutResponce{}, err
		}
	} else {
		return PutResponce{}, errors.New("not supported")
	}
	return PutResponce{}, nil
}

// DeletePriceLog delete a pricelogs
func (u *PriceLogController) DeletePriceLog(req DeleteRequest) (DeleteResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start DeletePriceLog", req)

	err := u.validateDeleteReq(req)
	if err != nil {
		u.logger.LogWrite(log.Error, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end DeletePriceLog:")
		return DeleteResponce{}, err
	}

	var inputModel = services.DeleteInputModel{UserID: req.UserID, GroupID: req.GroupID}
	res, err := u.service.DeletePriceLog(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Error, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end DeletePriceLog:")
		return DeleteResponce{}, err
	}

	u.logger.LogWriteWithMsgAndObj(log.Info, "end DeletePriceLog:", res)
	return DeleteResponce{}, nil
}

func (u *PriceLogController) validateGetReq(req GetRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.GroupID == "" || strings.TrimSpace(req.GroupID) == "" {
		return errors.New("GroupID is empty")
	}

	return nil
}

func (u *PriceLogController) validatePutReq(req PutRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.GroupID == "" || strings.TrimSpace(req.GroupID) == "" {
		return errors.New("GroupID is empty")
	}
	if !isPutMode(req.Mode) && !isUpdateMode(req.Mode) {
		return errors.New("Mode is not defined")
	}
	return nil
}

func isPutMode(mode string) bool {
	if mode == "put" {
		return true
	}
	return false
}

func isUpdateMode(mode string) bool {
	if mode == "update" {
		return true
	}
	return false
}

func (u *PriceLogController) validateDeleteReq(req DeleteRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.GroupID == "" || strings.TrimSpace(req.GroupID) == "" {
		return errors.New("GroupID is empty")
	}

	return nil
}
