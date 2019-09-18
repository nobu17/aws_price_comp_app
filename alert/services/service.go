package services

import (
	"alert/repositories"
	"common/log"
	"errors"
	"fmt"
)

// alertService serivce
type alertService struct {
	logger     log.LoggerImpl
	repository repositories.AlertImpl
}

// NewItemMasterService constructor
func NewItemMasterService(logger log.LoggerImpl, repository repositories.AlertImpl) ServiceImpl {
	return &alertService{
		logger:     logger,
		repository: repository,
	}
}

// GetAlertLog impl
func (u *alertService) GetAlertLog(req GetInputModel) (GetOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start GetAlertLog:", req)

	var input = repositories.GetRequest{UserID: req.UserID, MinAlertDate: req.MinAlertDate}
	res, err := u.repository.GetAlertLog(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end GetAlertLog")
		return GetOutputModel{}, err
	}
	list := make([]SendAlertLog, 0)
	for _, item := range res.SendAlertLogList {
		list = append(list, NewSendAlertLog(item.UserID, item.AlertDate, item.StoreType, item.ProductID, item.Price))
	}

	var output = GetOutputModel{SendAlertLogList: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetAlertLog:", output)

	return output, nil
}

func (u *alertService) PutAlertLog(req PutInputModel) error {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start PutAlertLog:", req)

	list := make([]repositories.SendAlertLog, 0)
	for _, item := range req.PutAlertLogList {
		list = append(list, repositories.NewSendAlertLog(item.UserID, item.AlertDate, item.StoreType, item.ProductID, item.Price))
	}
	input := repositories.PutRequest{PutAlertLogList: list}
	res, err := u.repository.PutAlertLog(input)
	if err != nil {
		return err
	}
	if len(list) != res.Wrote {
		return errors.New("some record is failed to write")
	}
	return nil
}
