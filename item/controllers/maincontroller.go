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
		u.logger.LogWrite(log.Error, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end GetItemMaster:")
		return GetResponce{}, err
	}

	var inputModel = services.InputModel{UserID: req.UserID, GroupID: req.GroupID}
	res, err := u.service.GetItemMasters(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Error, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end GetItemMaster:")
		return GetResponce{}, err
	}
	var list = make([]ItemMaster, 0)
	for _, item := range res.ItemMasters {
		list = append(list, NewItemMaster(item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName))
	}

	var respo = GetResponce{UserID: req.UserID, GroupID: req.GroupID, ItemMasters: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetItemMaster:", respo)
	return respo, nil
}

// PutItemMaster getitem
func (u *ItemMasterController) PutItemMaster(req PutRequest) (PutResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start PutItemMaster", req)

	err := u.validatePut(req)
	if err != nil {
		u.logger.LogWrite(log.Error, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end PutItemMaster:")
		return PutResponce{}, err
	}

	list := make([]services.ItemMaster, 0)
	for _, item := range req.ItemMasters {
		list = append(list, services.NewItemMaster(item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName))
	}
	var inputModel = services.PutInputModel{UserID: req.UserID, GroupID: req.GroupID, ItemMasters: list}
	_, err = u.service.PutItemMasters(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Error, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end PutItemMaster:")
		return PutResponce{}, err
	}

	u.logger.LogWrite(log.Info, "end PutItemMaster:")
	return PutResponce{}, nil
}

// DeleteItemMaster getitem
func (u *ItemMasterController) DeleteItemMaster(req DeleteRequest) (DeleteResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start DeleteItemMaster", req)

	err := u.validateDel(req)
	if err != nil {
		u.logger.LogWrite(log.Error, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end DeleteItemMaster:")
		return DeleteResponce{}, err
	}

	var inputModel = services.DeleteInputModel{UserID: req.UserID, GroupID: req.GroupID}
	_, err = u.service.DeleteItemMasters(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Error, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Error, "end DeleteItemMaster:")
		return DeleteResponce{}, err
	}

	u.logger.LogWrite(log.Info, "end DeleteItemMaster:")
	return DeleteResponce{}, nil
}

func (u *ItemMasterController) validatePut(req PutRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.GroupID == "" || strings.TrimSpace(req.GroupID) == "" {
		return errors.New("GroupID is empty")
	}
	if req.ItemMasters == nil || len(req.ItemMasters) == 0 {
		return errors.New("ItemMasters is empty")
	}
	if len(req.ItemMasters) == 0 {
		return errors.New("ItemMasters should be greater than 0")
	}
	if len(req.ItemMasters) > 20 {
		return errors.New("ItemMasters length should be less than 20")
	}
	dup := make(map[string]string)
	for _, item := range req.ItemMasters {
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
		// check duplicate id
		uniqueID := item.StoreType + "_" + item.ProductID
		_, existed := dup[uniqueID]
		if existed {
			return errors.New("Same Product ID set:" + uniqueID)
		}
		dup[uniqueID] = ""
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

func (u *ItemMasterController) validateDel(req DeleteRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.GroupID == "" || strings.TrimSpace(req.GroupID) == "" {
		return errors.New("GroupID is empty")
	}
	return nil
}
