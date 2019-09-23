package services

import (
	"common/log"
	"fmt"
	"item/repositories"
)

// productService serivce
type itemMasterService struct {
	logger     log.LoggerImpl
	repository repositories.ItemMasterImpl
}

// NewItemMasterService constructor
func NewItemMasterService(logger log.LoggerImpl, repository repositories.ItemMasterImpl) ServiceImpl {
	return &itemMasterService{
		logger:     logger,
		repository: repository,
	}
}

// GetItemMasters get item master.
func (u *itemMasterService) GetItemMasters(req InputModel) (OutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start GetItemMasters:", req)

	var input = repositories.Request{UserID: req.UserID, GroupID: req.GroupID}
	res, err := u.repository.GetItemMaster(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end GetItemMasters")
		return OutputModel{}, err
	}
	list := make([]ItemMaster, 0)
	for _, item := range res.ItemMasters {
		list = append(list, NewItemMaster(item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName))
	}

	var output = OutputModel{UserID: req.UserID, GroupID: req.GroupID, ItemMasters: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetItemMasters:", output)

	return output, nil
}

// PutItemMasters put item masters.
func (u *itemMasterService) PutItemMasters(req PutInputModel) (PutOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start PutItemMasters:", req)

	list := make([]repositories.ItemMaster, 0)
	for _, item := range req.ItemMasters {
		list = append(list, repositories.NewItemMaster(item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName))
	}
	var input = repositories.PutRequest{UserID: req.UserID, GroupID: req.GroupID, ItemMasters: list}
	_, err := u.repository.PutItemMaster(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutItemMasters")
		return PutOutputModel{}, err
	}

	u.logger.LogWrite(log.Info, "end PutItemMasters")

	return PutOutputModel{}, nil
}

// DeleteItemMasters del item masters.
func (u *itemMasterService) DeleteItemMasters(req DeleteInputModel) (DeleteOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start DeleteItemMasters:", req)

	var input = repositories.DeleteRequest{UserID: req.UserID, GroupID: req.GroupID}
	_, err := u.repository.DeleteItemMaster(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end DeleteItemMasters")
		return DeleteOutputModel{}, err
	}

	u.logger.LogWrite(log.Info, "end DeleteItemMasters")

	return DeleteOutputModel{}, nil
}
