package controllers

import (
	"common/log"
	"errors"
	"fmt"
	"notify/services"
)

// NotifyController controller.
type NotifyController struct {
	logger  log.LoggerImpl
	service services.ServiceImpl
}

// NewNotifyController constroctor
func NewNotifyController() NotifyController {
	container := Container{}
	ser := container.NewService()
	log := container.NewLogger()
	return NotifyController{
		logger:  log,
		service: ser,
	}
}

// SendNotify getitem
func (u *NotifyController) SendNotify(req PutRequest) (PutResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start NotifyController:SendNotify", req)

	err := u.validatePutReq(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end NotifyController:SendNotify:")
		return PutResponce{}, err
	}

	list := make([]services.ProductInfo, 0)
	for _, item := range req.ProductInfoList {
		list = append(list, services.NewProductInfo(item.ProductID, item.StoreType, item.Name, item.Price, item.ShippingFee))
	}
	user := services.NewUserInfo(req.UserInfo.UserID, req.UserInfo.Name, req.UserInfo.Mail)
	var input = services.PutInputModel{UserInfo: user, GroupID: req.GroupID, ProductInfoList: list}

	_, err = u.service.SendNotify(input)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end NotifyController:SendNotify:")
		return PutResponce{}, err
	}
	return PutResponce{}, nil
}

func (u *NotifyController) validatePutReq(req PutRequest) error {
	if req.UserInfo.UserID == "" {
		return errors.New("UserInfo.UserID is empty")
	}
	if req.UserInfo.Name == "" {
		return errors.New("UserInfo.Name is empty")
	}
	if req.UserInfo.Mail == "" {
		return errors.New("UserInfo.Mail is empty")
	}
	if req.GroupID == "" {
		return errors.New("UserInfo.GroupID is empty")
	}
	if len(req.ProductInfoList) < 1 {
		return errors.New("ProductInfoList is empty")
	}

	return nil
}
