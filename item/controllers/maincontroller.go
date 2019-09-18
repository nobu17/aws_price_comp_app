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
		list = append(list, NewItemMaster(item.GroupID, item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName))
	}

	var respo = GetResponce{ItemMasters: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetItemMaster:", respo)
	return respo, nil
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
