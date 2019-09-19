package services

import (
	"common/log"
	"fmt"
	"notify/repositories"
)

// alertService serivce
type sesService struct {
	logger     log.LoggerImpl
	repository repositories.SendNotifyImpl
}

// NewItemMasterService constructor
func NewItemMasterService(logger log.LoggerImpl, repository repositories.SendNotifyImpl) ServiceImpl {
	return &sesService{
		logger:     logger,
		repository: repository,
	}
}

// SendNotify impl
func (u *sesService) SendNotify(req PutInputModel) (PutOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start sesService:SendNotify", req)

	list := make([]repositories.ProductInfo, 0)
	for _, item := range req.ProductInfoList {
		list = append(list, repositories.NewProductInfo(item.ProductID, item.StoreType, item.Name, item.Price, item.ShippingFee))
	}

	user := repositories.NewUserInfo(req.UserInfo.UserID, req.UserInfo.Name, req.UserInfo.Mail)

	var input = repositories.PutRequest{UserInfo: user, GroupID: req.GroupID, ProductInfoList: list}
	_, err := u.repository.SendNotify(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end sesService:SendNotify")
		return PutOutputModel{}, err
	}

	u.logger.LogWrite(log.Info, "end sesService:SendNotify")

	return PutOutputModel{}, nil
}
