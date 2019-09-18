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

// GetItemMasters get itemmaster
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
		list = append(list, NewItemMaster(item.GroupID, item.ProductID, item.StoreType, item.ThretholdPrice, item.ItemName))
	}

	var output = OutputModel{ItemMasters: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetItemMasters:", output)

	return output, nil
}
