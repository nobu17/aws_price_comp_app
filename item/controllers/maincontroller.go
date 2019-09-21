package controllers

import (
	"common/log"
	"errors"
	"fmt"
	"item/services"
	"strings"
)

// ItemMasterController controller
type ItemMasterController struct {
	logger  log.LoggerImpl
	service services.ServiceImpl
}

// NewItemMasterController constroctor
func NewItemMasterController() ItemMasterController {
	container := Container{}
	ser := container.NewService()
	log := container.NewLogger()
	return ItemMasterController{
		logger:  log,
		service: ser,
	}
}

// GetItemMaster getitem
func (u *ItemMasterController) GetItemMaster(req GetRequest) (GetResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start GetItemMaster", req)

	err := u.validate(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end GetItemMaster:")
		return GetResponce{}, err
	}

	var inputModel = services.InputModel{UserID: req.UserID, GroupID: req.GroupID}
	res, err := u.service.GetItemMasters(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end GetItemMaster:")
		return GetResponce{}, err
	}
	var list = make([]ItemMaster, 0)
	for _, item := range res.ItemMasters {
		list = append(list, NewItemMaster(item.UserID, item.GroupID, item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName))
	}

	var respo = GetResponce{ItemMasters: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetItemMaster:", respo)
	return respo, nil
}

// PutItemMaster getitem
func (u *ItemMasterController) PutItemMaster(req PutRequest) (PutResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start PutItemMaster", req)

	err := u.validatePut(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutItemMaster:")
		return PutResponce{}, err
	}

	list := make([]services.ItemMaster, 0)
	for _, item := range req.ItemMasters {
		list = append(list, services.NewItemMaster(item.UserID, item.GroupID, item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName))
	}
	var inputModel = services.PutInputModel{ItemMasters: list}
	_, err = u.service.PutItemMasters(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutItemMaster:")
		return PutResponce{}, err
	}

	u.logger.LogWrite(log.Info, "end PutItemMaster:")
	return PutResponce{}, nil
}

func (u *ItemMasterController) validatePut(req PutRequest) error {
	if req.ItemMasters == nil || len(req.ItemMasters) == 0 {
		return errors.New("ItemMasters is empty")
	}
	for _, item := range req.ItemMasters {
		if (item.UserID == "") || strings.TrimSpace(item.UserID) == "" {
			return errors.New("UserID is empty")
		}
		if (item.GroupID == "") || strings.TrimSpace(item.GroupID) == "" {
			return errors.New("GroupID is empty")
		}
		if (item.ProductID == "") || strings.TrimSpace(item.ProductID) == "" {
			return errors.New("ProductID is empty")
		}
		if (item.StoreType == "") || strings.TrimSpace(item.StoreType) == "" {
			return errors.New("StoreType is empty")
		}
		if item.ThretholdPrice < 1 {
			return errors.New("ThretholdPrice is less than 1")
		}
		if (item.ItemName == "") || strings.TrimSpace(item.ItemName) == "" {
			return errors.New("ItemName is empty")
		}
	}
	return nil
}

func (u *ItemMasterController) validate(req GetRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.GroupID == "" || strings.TrimSpace(req.GroupID) == "" {
		return errors.New("GroupID is empty")
	}
	return nil
}
