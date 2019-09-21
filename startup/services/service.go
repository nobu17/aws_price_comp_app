package services

import (
	"common/log"
	"fmt"
	"startup/repositories"
)

// observeService service interface
type observeService struct {
	logger         log.LoggerImpl
	userRepository repositories.UserInfoImpl
	sendRepository repositories.SendObservImpl
}

// NewObserveService constructor
func NewObserveService(logger log.LoggerImpl, userRepository repositories.UserInfoImpl, sendRepository repositories.SendObservImpl) ServiceImpl {
	return &observeService{logger: logger, userRepository: userRepository, sendRepository: sendRepository}
}

func (u *observeService) StartObserve(req InputModel) (OutputModel, error) {
	// get user
	u.logger.LogWriteWithMsgAndObj(log.Info, "start observeService:StartObserve:", req)
	userParam := repositories.NewGetRequest(req.UserID, req.Password)
	res, err := u.userRepository.GetUserInfo(userParam)
	if err != nil {
		u.logger.LogWrite(log.Error, "get user is failed:"+fmt.Sprint(err))
		return OutputModel{}, err
	}
	u.logger.LogWriteWithMsgAndObj(log.Error, "observeService:GetUserInfo:", res)

	// send que
	var sendParam = repositories.SendRequest{UserInfo: res.UserInfo, ItemGroupList: res.ItemGroupList}
	err = u.sendRepository.SendObservRequest(sendParam)
	if err != nil {
		u.logger.LogWrite(log.Error, "SendObservRequest is failed:"+fmt.Sprint(err))
	}
	u.logger.LogWrite(log.Info, "end observeService:StartObserve:")
	return OutputModel{}, err
}
