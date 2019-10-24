package services

import (
	"common/log"
	"common/util"
	"fmt"
	"pricelog/repositories"
)

// productService serivce
type priceLogMasterService struct {
	logger     log.LoggerImpl
	repository repositories.PriceLogImpl
}

// NewPriceLogMasterService constructor
func NewPriceLogMasterService(logger log.LoggerImpl, repository repositories.PriceLogImpl) ServiceImpl {
	return &priceLogMasterService{
		logger:     logger,
		repository: repository,
	}
}

// GetPriceLog getprice log list
func (u *priceLogMasterService) GetPriceLog(req GetInputModel) (GetOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start GetPriceLog:", req)

	var input = repositories.GetRequest{UserID: req.UserID, GroupID: req.GroupID}
	res, err := u.repository.GetPriceLogs(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end GetPriceLog")
		return GetOutputModel{}, err
	}

	list := make([]PriceLog, 0)
	for _, item := range res.PriceLogList {
		list = append(list, PriceLog{StoreType: item.StoreType, ItemID: item.ItemID, Price: item.Price, LastModifiedDatetime: item.LastModifiedDatetime})
	}

	var output = GetOutputModel{UserID: req.UserID, GroupID: req.GroupID, PriceLogList: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetPriceLog:", output)

	return output, nil
}

// PutPriceLog put price log list
func (u *priceLogMasterService) PutPriceLog(req PutInputModel) (PutOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start PutPriceLog:", req)

	list := make([]repositories.PriceLog, 0)
	for _, item := range req.PriceLogList {
		list = append(list, repositories.PriceLog{StoreType: item.StoreType, ItemID: item.ItemID, Price: item.Price, LastModifiedDatetime: util.GetNowJSTTimeStr()})
	}
	var input = repositories.PutRequest{UserID: req.UserID, GroupID: req.GroupID, PriceLogList: list}
	_, err := u.repository.PutPriceLogs(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutPriceLog")
		return PutOutputModel{}, err
	}

	u.logger.LogWrite(log.Info, "end PutPriceLog")

	return PutOutputModel{}, nil
}

// UpdatePriceLog update price log list
func (u *priceLogMasterService) UpdatePriceLog(req UpdateInputModel) (UpdateOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start UpdatePriceLog:", req)

	// get current price log
	var input = repositories.GetRequest{UserID: req.UserID, GroupID: req.GroupID}
	res, err := u.repository.GetPriceLogs(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end UpdatePriceLog")
		return UpdateOutputModel{}, err
	}
	// compare and if detected more cheaper price, then update
	list := make([]repositories.PriceLog, 0)
	for _, item := range res.PriceLogList {
		isUpdated := false
		for _, cItem := range req.PriceLogList {
			if item.StoreType == cItem.StoreType && item.ItemID == cItem.ItemID {
				if item.Price > cItem.Price {
					list = append(list, repositories.PriceLog{StoreType: cItem.StoreType, ItemID: cItem.ItemID, Price: cItem.Price, LastModifiedDatetime: util.GetNowJSTTimeStr()})
				} else {
					list = append(list, repositories.PriceLog{StoreType: item.StoreType, ItemID: item.ItemID, Price: item.Price, LastModifiedDatetime: item.LastModifiedDatetime})
				}
				isUpdated = true
			}
		}
		if !isUpdated {
			list = append(list, repositories.PriceLog{StoreType: item.StoreType, ItemID: item.ItemID, Price: item.Price, LastModifiedDatetime: item.LastModifiedDatetime})
		}
	}
	// add without it is already added
	for _, item := range req.PriceLogList {
		isAlreadyAdd := false
		for _, cItem := range list {
			if item.StoreType == cItem.StoreType && item.ItemID == cItem.ItemID {
				isAlreadyAdd = true
				break
			}
		}
		if !isAlreadyAdd {
			list = append(list, repositories.PriceLog{StoreType: item.StoreType, ItemID: item.ItemID, Price: item.Price, LastModifiedDatetime: util.GetNowJSTTimeStr()})
		}
	}

	var putinput = repositories.PutRequest{UserID: req.UserID, GroupID: req.GroupID, PriceLogList: list}
	_, err = u.repository.PutPriceLogs(putinput)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end PutPriceLog")
		return UpdateOutputModel{}, err
	}
	return UpdateOutputModel{}, nil
}

// DeletePriceLog put price log list
func (u *priceLogMasterService) DeletePriceLog(req DeleteInputModel) (DeleteOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start DeletePriceLog:", req)

	var input = repositories.DeleteRequest{UserID: req.UserID, GroupID: req.GroupID}
	_, err := u.repository.DeletePriceLogs(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end DeletePriceLog")
		return DeleteOutputModel{}, err
	}

	u.logger.LogWrite(log.Info, "end DeletePriceLog")

	return DeleteOutputModel{}, nil
}
